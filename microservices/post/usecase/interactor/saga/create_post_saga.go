package saga

import (
	"context"
	"errors"
	"time"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/pb"
	"github.com/ezio1119/fishapp-post/usecase/repo"
	"github.com/looplab/fsm"
	"google.golang.org/protobuf/encoding/protojson"
)

type createPostSagaState struct {
	sagaID       string
	sagaType     string
	currentState string
	post         *pb.Post
	createdAt    time.Time
	updatedAt    time.Time
}

func NewCreatePostSagaState(p *pb.Post, state, sagaID string) (*createPostSagaState, error) {
	now := time.Now()
	return &createPostSagaState{
		sagaType:     "CreatePostSaga",
		post:         p,
		sagaID:       sagaID,
		currentState: state,
		createdAt:    now,
		updatedAt:    now,
	}, nil
}

func (s *createPostSagaState) convSagaInstance(state string) (*models.SagaInstance, error) {
	jsonPost, err := protojson.Marshal(s.post)
	if err != nil {
		return nil, err
	}
	return &models.SagaInstance{
		ID:           s.sagaID,
		SagaType:     s.sagaType,
		SagaData:     jsonPost,
		CurrentState: state,
		CreatedAt:    s.createdAt,
		UpdatedAt:    s.updatedAt,
	}, nil
}

type CreatePostSagaManager struct {
	FSM              *fsm.FSM
	state            *createPostSagaState
	outboxRepo       repo.OutboxRepo
	postRepo         repo.PostRepo
	sagaInstanceRepo repo.SagaInstanceRepo
	transactionRepo  repo.TransactionRepo
}

func InitCreatePostSagaManager(
	or repo.OutboxRepo,
	pr repo.PostRepo,
	sr repo.SagaInstanceRepo,
	tr repo.TransactionRepo,
) *CreatePostSagaManager {
	return &CreatePostSagaManager{
		outboxRepo:       or,
		postRepo:         pr,
		sagaInstanceRepo: sr,
		transactionRepo:  tr,
	}
}

func (m *CreatePostSagaManager) NewCreatePostSagaManager(state *createPostSagaState) *CreatePostSagaManager {

	m.state = state

	m.FSM = fsm.NewFSM(
		"init",
		fsm.Events{
			// {Name: "UploadImage", Src: []string{"Init"}, Dst: "UploadingImage"},
			{Name: "CreateRoom", Src: []string{"init"}, Dst: "CreatingRoom"},
			{Name: "RejectPost", Src: []string{"CreatingRoom"}, Dst: "PostRejected"},
			{Name: "ApprovePost", Src: []string{"CreatingRoom"}, Dst: "PostApproved"},
		},
		fsm.Callbacks{
			// "UploadImage": func(e *fsm.Event) { s.uploadImage(e) },
			"CreateRoom":  func(e *fsm.Event) { m.createRoom(e) },
			"RejectPost":  func(e *fsm.Event) { m.rejectPost(e) },
			"ApprovePost": func(e *fsm.Event) { m.approvePost(e) },
		},
	)

	m.FSM.SetState(m.state.currentState)

	return m
}

func (m *CreatePostSagaManager) createRoom(e *fsm.Event) {
	ctx, ok := e.Args[0].(context.Context)
	if !ok {
		e.Cancel(errors.New("missing context"))
		return
	}
	// 遷移先のステートを入れる
	sagaIn, err := m.state.convSagaInstance(e.Dst)
	if err != nil {
		e.Cancel(err)
		return
	}

	event, err := newCreateRoomEvent(&pb.CreateRoom{
		SagaId: m.state.sagaID,
		PostId: m.state.post.Id,
		UserId: m.state.post.UserId,
	})
	if err != nil {
		e.Cancel(err)
		return
	}
	// sagaインスタンスの永続化とイベントの発行を同じトランザクション内でやる
	ctx, err = m.transactionRepo.BeginTx(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}

	defer func() {
		if recover() != nil {
			m.transactionRepo.Roolback(ctx)
		}
	}()

	if err := m.outboxRepo.CreateOutbox(ctx, event); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	if err := m.sagaInstanceRepo.CreateSagaInstance(ctx, sagaIn); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	ctx, err = m.transactionRepo.Commit(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}
	// e.Cancel(errors.New("errorおきたよ"))
	m.state.currentState = e.Dst
}

func (m *CreatePostSagaManager) rejectPost(e *fsm.Event) {
	ctx, ok := e.Args[0].(context.Context)
	if !ok {
		e.Cancel(errors.New("missing context"))
		return
	}

	errMsg, ok := e.Args[1].(string)
	if !ok {
		e.Cancel(errors.New("missing error message"))
		return
	}

	event, err := newPostRejectedEvent(m.state.post, m.state.sagaID, errMsg)
	if err != nil {
		e.Cancel(err)
		return
	}

	jsonPost, err := protojson.Marshal(m.state.post)
	if err != nil {
		e.Cancel(err)
		return
	}

	now := time.Now()
	sagaIn := &models.SagaInstance{
		ID:           m.state.sagaID,
		SagaType:     m.state.sagaType,
		SagaData:     jsonPost,
		CurrentState: e.Dst,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	// createPostSagaFailedとサガイベントも発行する
	ctx, err = m.transactionRepo.BeginTx(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}

	defer func() {
		if recover() != nil {
			m.transactionRepo.Roolback(ctx)
		}
	}()

	if err := m.postRepo.DeletePost(ctx, m.state.post.Id); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	if err := m.outboxRepo.CreateOutbox(ctx, event); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	if err := m.sagaInstanceRepo.UpdateSagaInstance(ctx, sagaIn); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	ctx, err = m.transactionRepo.Commit(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}
}

func (m *CreatePostSagaManager) approvePost(e *fsm.Event) {
	ctx, ok := e.Args[0].(context.Context)
	if !ok {
		e.Cancel(errors.New("missing context"))
		return
	}

	event, err := newPostApprovedEvent(m.state.post, m.state.sagaID)
	if err != nil {
		e.Cancel(err)
		return
	}

	jsonPost, err := protojson.Marshal(m.state.post)
	if err != nil {
		e.Cancel(err)
		return
	}

	sagaIn := &models.SagaInstance{
		ID:           m.state.sagaID,
		SagaType:     m.state.sagaType,
		SagaData:     jsonPost,
		CurrentState: e.Dst,
		UpdatedAt:    time.Now(),
	}

	ctx, err = m.transactionRepo.BeginTx(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}

	defer func() {
		if recover() != nil {
			m.transactionRepo.Roolback(ctx)
		}
	}()

	if err := m.outboxRepo.CreateOutbox(ctx, event); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	if err := m.sagaInstanceRepo.UpdateSagaInstance(ctx, sagaIn); err != nil {
		m.transactionRepo.Roolback(ctx)
		e.Cancel(err)
		return
	}

	ctx, err = m.transactionRepo.Commit(ctx)
	if err != nil {
		e.Cancel(err)
		return
	}

	m.state.currentState = e.Dst
}
