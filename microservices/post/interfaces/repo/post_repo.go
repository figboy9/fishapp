package repo

import (
	"context"
	"fmt"
	"log"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/usecase/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postRepo struct {
	SqlHandler
}

func NewPostRepo(h SqlHandler) repo.PostRepo {
	return &postRepo{h}
}

func (r *postRepo) fetchPosts(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	result := make([]*models.Post, 0)
	for rows.Next() {
		p := new(models.Post)
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.FishingSpotTypeID,
			&p.PrefectureID,
			&p.MeetingPlaceID,
			&p.MeetingAt,
			&p.MaxApply,
			&p.UserID,
			&p.UpdatedAt,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}

	return result, nil
}

func (r *postRepo) fetchPostsFishTypes(ctx context.Context, query string, args ...interface{}) ([]*models.PostsFishType, error) {
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	result := make([]*models.PostsFishType, 0)
	for rows.Next() {
		f := new(models.PostsFishType)
		if err := rows.Scan(
			&f.ID,
			&f.PostID,
			&f.FishTypeID,
			&f.UpdatedAt,
			&f.CreatedAt,
		); err != nil {
			return nil, err
		}
		result = append(result, f)
	}

	return result, nil
}

func (r *postRepo) fillPostWithFishTypeIDs(ctx context.Context, p *models.Post) error {
	query := `SELECT id, post_id, fish_type_id, updated_at, created_at
           	FROM posts_fish_types
						WHERE post_id = ?`

	fishes, err := r.fetchPostsFishTypes(ctx, query, p.ID)
	if err != nil {
		return err
	}
	for _, f := range fishes {
		p.PostsFishTypes = append(p.PostsFishTypes, f)
	}

	return nil
}

func (r *postRepo) fillListPostsWithFishTypes(ctx context.Context, posts []*models.Post) error {
	query := `SELECT id, post_id, fish_type_id, updated_at, created_at
            FROM posts_fish_types
            WHERE post_id IN(?` + strings.Repeat(",?", len(posts)-1) + ")"

	args := make([]interface{}, len(posts))
	for i, p := range posts {
		args[i] = p.ID
	}

	fishes, err := r.fetchPostsFishTypes(ctx, query, args...)
	if err != nil {
		return err
	}

	for _, p := range posts {
		for _, f := range fishes {
			if p.ID == f.PostID {
				p.PostsFishTypes = append(p.PostsFishTypes, f)
			}
		}
	}

	return nil
}

func (r *postRepo) batchCreatePostsFishTypes(ctx context.Context, p *models.Post) error {
	query := `INSERT INTO posts_fish_types(post_id, fish_type_id, created_at, updated_at)
						VALUES (?, ?, ?, ?)` + strings.Repeat(", (?, ?, ?, ?)", len(p.PostsFishTypes)-1)

	for _, f := range p.PostsFishTypes {
		f.PostID = p.ID
	}

	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	args := []interface{}{}
	for _, f := range p.PostsFishTypes {
		args = append(args, f.PostID, f.FishTypeID, f.CreatedAt, f.UpdatedAt)
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(rowCnt) != len(p.PostsFishTypes) {
		return fmt.Errorf("expected %d row affected, got %d rows affected", len(p.PostsFishTypes), rowCnt)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	for i, f := range p.PostsFishTypes {
		f.ID = lastID - int64(len(p.PostsFishTypes)) + int64(i) + 2
	}

	return nil
}

func (r *postRepo) deletePostsFishTypesByPostID(ctx context.Context, pID int64) error {
	query := "DELETE FROM posts_fish_types WHERE post_id = ?"
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, pID); err != nil {
		return err
	}

	return nil
}

func (r *postRepo) CreatePost(ctx context.Context, p *models.Post) error {
	query := `INSERT posts SET title=?, content=?, fishing_spot_type_id=?, prefecture_id=?, meeting_place_id=?, meeting_at=?, max_apply=?, user_id=?, updated_at=?, created_at=?`
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.Title, p.Content, p.FishingSpotTypeID, p.PrefectureID, p.MeetingPlaceID, p.MeetingAt, p.MaxApply, p.UserID, p.UpdatedAt, p.CreatedAt)
	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(rowCnt) != 1 {
		return fmt.Errorf("expected %d row affected, got %d rows affected", 1, rowCnt)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = lastID

	if err := r.batchCreatePostsFishTypes(ctx, p); err != nil {
		return err
	}

	return nil
}

func (r *postRepo) GetPostByID(ctx context.Context, id int64) (*models.Post, error) {
	query := `SELECT id, title, content, fishing_spot_type_id, prefecture_id, meeting_place_id, meeting_at, max_apply, user_id, updated_at, created_at
            FROM posts
            WHERE id = ?`

	list, err := r.fetchPosts(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		return nil, status.Errorf(codes.NotFound, "post with id='%d' is not found", id)
	}

	if err := r.fillPostWithFishTypeIDs(ctx, list[0]); err != nil {
		return nil, err
	}

	return list[0], nil
}

func (r *postRepo) ListPosts(ctx context.Context, p *models.Post, num int64, cursor int64, f *models.PostFilter) ([]*models.Post, error) {
	sq := sq.Select("id, title, content, fishing_spot_type_id, prefecture_id, meeting_place_id, meeting_at, max_apply, user_id, updated_at, created_at").
		From("posts").
		GroupBy("posts.id").
		Limit(uint64(num))

	if p.FishingSpotTypeID != 0 {
		sq = sq.Where("fishing_spot_type_id = ?", p.FishingSpotTypeID)
	}

	if p.PrefectureID != 0 {
		sq = sq.Where("prefecture_id = ?", p.PrefectureID)
	}

	if p.UserID != 0 {
		sq = sq.Where("user_id = ?", p.UserID)
	}

	if f.CanApply {
		sq = sq.LeftJoin("apply_posts ON posts.id = apply_posts.post_id").
			Having("count(apply_posts.id) < posts.max_apply")
	}

	if f.FishTypeIDs != nil {
		sq = sq.Join("posts_fish_types ON posts.id = posts_fish_types.post_id").
			Where("posts_fish_types.fish_type_id IN(?)", f.FishTypeIDs).
			Having("count(posts_fish_types.fish_type_id) = ?", len(f.FishTypeIDs))
	}

	if !f.MeetingAtFrom.IsZero() && !f.MeetingAtTo.IsZero() {
		sq = sq.Where("meeting_at BETWEEN ? AND ?", f.MeetingAtFrom, f.MeetingAtTo)
	}

	if cursor != 0 {
		switch f.SortBy {
		case models.SortByID:

			if f.OrderBy == models.OrderByAsc {
				sq = sq.Where("posts.id > ?", cursor).
					OrderBy("id asc")
			}

			if f.OrderBy == models.OrderByDesc {
				sq = sq.Where("posts.id < ?", cursor).
					OrderBy("id desc")
			}
		// meeting_atはユニークではないため、同じ値の場合を考えidでも絞り込む
		case models.SortByMeetingAt:

			p, err := r.GetPostByID(ctx, cursor)
			if err != nil {
				return nil, err
			}

			switch f.OrderBy {
			case models.OrderByAsc:

				sq = sq.Where("meeting_at >= ?", p.MeetingAt).
					Where("meeting_at > ? or posts.id > ?", p.MeetingAt, cursor).
					OrderBy("meeting_at asc, id asc")

			case models.OrderByDesc:

				sq = sq.Where("meeting_at <= ?", p.MeetingAt).
					Where("meeting_at < ? or posts.id < ?", p.MeetingAt, cursor).
					OrderBy("meeting_at desc, id desc")

			}
		}
	}

	if cursor == 0 {
		sq = sq.OrderBy(fmt.Sprintf("%s %s", f.SortBy, f.OrderBy))
	}

	query, args, err := sq.ToSql()
	fmt.Println(query)
	if err != nil {
		return nil, err
	}

	posts, err := r.fetchPosts(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	if len(posts) != 0 {
		if err := r.fillListPostsWithFishTypes(ctx, posts); err != nil {
			return nil, err
		}
	}

	return posts, nil
}

func (r *postRepo) UpdatePost(ctx context.Context, p *models.Post) error {
	query := `UPDATE posts SET title=?, content=?, fishing_spot_type_id=?, prefecture_id=?, meeting_place_id=?, meeting_at=?, max_apply=?, updated_at=?
						WHERE id = ?`

	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.Title, p.Content, p.FishingSpotTypeID, p.PrefectureID, p.MeetingPlaceID, p.MeetingAt, p.MaxApply, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowCnt != 1 {
		return fmt.Errorf("expected %d row affected, got %d rows affected", 1, rowCnt)
	}

	if err := r.deletePostsFishTypesByPostID(ctx, p.ID); err != nil {
		return err
	}

	if err := r.batchCreatePostsFishTypes(ctx, p); err != nil {
		return err
	}

	return nil
}

func (r *postRepo) DeletePost(ctx context.Context, id int64) error {
	query := "DELETE FROM posts WHERE id = ?"
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {

		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt != 1 {
		return fmt.Errorf("expected %d row affected, got %d rows affected", 1, rowCnt)
	}

	return nil
}
