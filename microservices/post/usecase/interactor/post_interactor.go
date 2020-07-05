package interactor

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/ezio1119/fishapp-post/conf"
	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/usecase/interactor/saga"
	"github.com/ezio1119/fishapp-post/usecase/repo"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PostInteractor interface {
	GetPost(ctx context.Context, id int64) (*models.Post, error)
	ListPosts(ctx context.Context, p *models.Post, pageSize int64, pageToken string, filter *models.PostFilter) ([]*models.Post, string, error)
	CreatePost(ctx context.Context, p *models.Post, imageBufs []*bytes.Buffer) (string, error)
	UpdatePost(ctx context.Context, p *models.Post, imageBufs []*bytes.Buffer, deleteImageIDs []int64) error
	DeletePost(ctx context.Context, id int64) error

	GetApplyPost(ctx context.Context, id int64) (*models.ApplyPost, error)
	ListApplyPosts(ctx context.Context, applyPost *models.ApplyPost) ([]*models.ApplyPost, error)
	BatchGetApplyPostsByPostIDs(ctx context.Context, postIDs []int64) ([]*models.ApplyPost, error)
	CreateApplyPost(ctx context.Context, applyPost *models.ApplyPost) error
	DeleteApplyPost(ctx context.Context, id int64) error
}

type postInteractor struct {
	postRepo              repo.PostRepo
	imageRepo             repo.ImageRepo
	applyPostRepo         repo.ApplyPostRepo
	transactionRepo       repo.TransactionRepo
	outboxRepo            repo.OutboxRepo
	createPostSagaManager *saga.CreatePostSagaManager
	ctxTimeout            time.Duration
}

func NewPostInteractor(
	pr repo.PostRepo,
	ir repo.ImageRepo,
	ar repo.ApplyPostRepo,
	tr repo.TransactionRepo,
	or repo.OutboxRepo,
	sm *saga.CreatePostSagaManager,
	timeout time.Duration,
) PostInteractor {
	return &postInteractor{pr, ir, ar, tr, or, sm, timeout}
}

func (i *postInteractor) GetPost(ctx context.Context, id int64) (*models.Post, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	p, err := i.postRepo.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (i *postInteractor) ListPosts(ctx context.Context, p *models.Post, pageSize int64, pageToken string, f *models.PostFilter) ([]*models.Post, string, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	fmt.Printf("post: %#v\npageSize: %#v\npageToken: %#v\nPostFilter: %#v\n", p, pageSize, pageToken, f)
	if pageSize == 0 {
		pageSize = conf.C.Sv.DefaultPageSize
	}

	pageSize++
	var cursor int64
	if pageToken != "" {
		var err error
		cursor, err = extractIDFromPageToken(pageToken)
		if err != nil {
			return nil, "", err
		}
	}

	list, err := i.postRepo.ListPosts(ctx, p, pageSize, cursor, f)
	if err != nil {
		return nil, "", err
	}
	nextToken := ""
	if len(list) == int(pageSize) {
		list = list[:pageSize-1]
		nextToken = genPageTokenFromID(list[len(list)-1].ID)
	}

	return list, nextToken, nil
}

func (i *postInteractor) CreatePost(ctx context.Context, p *models.Post, imageBufs []*bytes.Buffer) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now

	for _, f := range p.PostsFishTypes {
		f.CreatedAt = now
		f.UpdatedAt = now
	}

	ctx, err := i.transactionRepo.BeginTx(ctx)
	if err != nil {
		return "", err
	}

	defer func() {
		if recover() != nil {
			i.transactionRepo.Roolback(ctx)
		}
	}()

	// use tx inside
	if err := i.postRepo.CreatePost(ctx, p); err != nil {
		i.transactionRepo.Roolback(ctx)
		return "", err
	}

	ctx, err = i.transactionRepo.Commit(ctx)
	if err != nil {
		return "", err
	}

	if len(imageBufs) != 0 {
		if err := i.imageRepo.BatchCreateImages(ctx, p.ID, imageBufs); err != nil {
			if err := i.postRepo.DeletePost(ctx, p.ID); err != nil {
				return "", err
			}
			return "", err
		}
	}

	sagaID := uuid.New().String()
	pProto, err := convPostProto(p)
	if err != nil {
		return "", err
	}

	state, err := saga.NewCreatePostSagaState(pProto, "init", sagaID)
	if err != nil {
		return "", err
	}

	s := i.createPostSagaManager.NewCreatePostSagaManager(state)

	// 非同期
	if err := s.FSM.Event("CreateRoom", ctx); err != nil {
		return "", err
	}

	return sagaID, nil
}

// imageBufsは新しいイメージ、deleteImageIDsは消すimageのID
func (i *postInteractor) UpdatePost(ctx context.Context, p *models.Post, imageBufs []*bytes.Buffer, dltImageIDs []int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	fmt.Println(dltImageIDs)

	now := time.Now()

	oldP, err := i.postRepo.GetPostByID(ctx, p.ID)
	if err != nil {
		return err
	}

	// 	// 完全なデータにする
	p.UserID = oldP.UserID
	p.CreatedAt = oldP.CreatedAt
	p.UpdatedAt = now
	for _, f := range p.PostsFishTypes {
		f.CreatedAt = now
		f.UpdatedAt = now
	}

	ctx, err = i.transactionRepo.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if recover() != nil {
			i.transactionRepo.Roolback(ctx)
		}
	}()

	cnt, err := i.applyPostRepo.CountApplyPostsByPostID(ctx, p.ID)
	if err != nil {
		return err
	}

	if cnt > p.MaxApply {
		return status.Errorf(codes.FailedPrecondition, "got max_apply is %d but already have %d apply", p.MaxApply, cnt)
	}

	if err := i.postRepo.UpdatePost(ctx, p); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	ctx, err = i.transactionRepo.Commit(ctx)
	if err != nil {
		return err
	}

	if len(dltImageIDs) != 0 {
		if err := i.imageRepo.BatchDeleteImages(ctx, dltImageIDs); err != nil {
			if err := i.postRepo.UpdatePost(ctx, oldP); err != nil {
				return err
			}
			return err
		}

	}

	if len(imageBufs) != 0 {
		if err := i.imageRepo.BatchCreateImages(ctx, p.ID, imageBufs); err != nil {
			if err := i.postRepo.UpdatePost(ctx, oldP); err != nil {
				return err
			}
			return err
		}
	}

	return nil
}

func (i *postInteractor) DeletePost(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	p, err := i.postRepo.GetPostByID(ctx, id)
	if err != nil {
		return err
	}

	event, err := newPostDeletedEvent(p)
	if err != nil {
		return err
	}

	ctx, err = i.transactionRepo.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if recover() != nil {
			i.transactionRepo.Roolback(ctx)
		}
	}()

	if err := i.postRepo.DeletePost(ctx, id); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	if err := i.outboxRepo.CreateOutbox(ctx, event); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	ctx, err = i.transactionRepo.Commit(ctx)
	if err != nil {
		return err
	}

	if err := i.imageRepo.DeleteImagesByPostID(ctx, id); err != nil {
		return err
	}

	return nil
}

func (i *postInteractor) GetApplyPost(ctx context.Context, id int64) (*models.ApplyPost, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.applyPostRepo.GetApplyPostByID(ctx, id)
}

func (i *postInteractor) ListApplyPosts(ctx context.Context, a *models.ApplyPost) ([]*models.ApplyPost, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	if a.UserID != 0 {
		return i.applyPostRepo.ListApplyPostsByUserID(ctx, a.UserID)
	}
	if a.PostID != 0 {
		return i.applyPostRepo.ListApplyPostsByPostID(ctx, a.PostID)
	}
	return nil, nil
}

func (i *postInteractor) BatchGetApplyPostsByPostIDs(ctx context.Context, postIDs []int64) ([]*models.ApplyPost, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.applyPostRepo.BatchGetApplyPostsByPostIDs(ctx, postIDs)
}

func (i *postInteractor) CreateApplyPost(ctx context.Context, a *models.ApplyPost) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	now := time.Now()
	a.CreatedAt = now
	a.UpdatedAt = now

	ctx, err := i.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil
	}

	defer func() {
		if recover() != nil {
			i.transactionRepo.Roolback(ctx)
		}
	}()

	cnt, err := i.applyPostRepo.CountApplyPostsByPostID(ctx, a.PostID)
	if err != nil {
		return err
	}

	p, err := i.postRepo.GetPostByID(ctx, a.PostID)
	if err != nil {
		return err
	}

	if p.MaxApply <= cnt {
		return status.Error(codes.FailedPrecondition, "already reached max_apply limit")
	}

	if err := i.applyPostRepo.CreateApplyPost(ctx, a); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	event, err := newApplyPostCreatedEvent(a)
	if err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	if err := i.outboxRepo.CreateOutbox(ctx, event); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	ctx, err = i.transactionRepo.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (i *postInteractor) DeleteApplyPost(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	a, err := i.applyPostRepo.GetApplyPostByID(ctx, id)
	if err != nil {
		return err
	}

	event, err := newApplyPostDeletedEvent(a)
	if err != nil {
		return err
	}

	ctx, err = i.transactionRepo.BeginTx(ctx)
	if err != nil {
		return nil
	}

	defer func() {
		if recover() != nil {
			i.transactionRepo.Roolback(ctx)
		}
	}()

	if err := i.applyPostRepo.DeleteApplyPost(ctx, a.ID); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	if err := i.outboxRepo.CreateOutbox(ctx, event); err != nil {
		i.transactionRepo.Roolback(ctx)
		return err
	}

	ctx, err = i.transactionRepo.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
