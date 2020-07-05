package repo

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/usecase/repo"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type applyPostRepo struct {
	SqlHandler
}

func NewApplyPostRepo(h SqlHandler) repo.ApplyPostRepo {
	return &applyPostRepo{h}
}

func (r *applyPostRepo) fetchApplyPosts(ctx context.Context, query string, args ...interface{}) ([]*models.ApplyPost, error) {
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

	result := make([]*models.ApplyPost, 0)
	for rows.Next() {
		a := new(models.ApplyPost)
		err = rows.Scan(
			&a.ID,
			&a.PostID,
			&a.UserID,
			&a.UpdatedAt,
			&a.CreatedAt,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (r *applyPostRepo) GetApplyPostByID(ctx context.Context, id int64) (*models.ApplyPost, error) {
	query := `SELECT id, post_id, user_id, updated_at, created_at
                        FROM apply_posts
                        WHERE id = ?`
	list, err := r.fetchApplyPosts(ctx, query, id)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, status.Errorf(codes.NotFound, "apply_post with id='%d' is not found", id)
	}
	return list[0], nil
}

func (r *applyPostRepo) ListApplyPostsByUserID(ctx context.Context, uID int64) ([]*models.ApplyPost, error) {
	query := `SELECT id, post_id, user_id, updated_at, created_at
                        FROM apply_posts
                        WHERE user_id = ?`
	return r.fetchApplyPosts(ctx, query, uID)
}

func (r *applyPostRepo) ListApplyPostsByPostID(ctx context.Context, pID int64) ([]*models.ApplyPost, error) {
	query := `SELECT id, post_id, user_id, updated_at, created_at
                        FROM apply_posts
                        WHERE post_id = ?`
	return r.fetchApplyPosts(ctx, query, pID)
}

func (r *applyPostRepo) BatchGetApplyPostsByPostIDs(ctx context.Context, pIDs []int64) ([]*models.ApplyPost, error) {
	query := `SELECT id, post_id, user_id, updated_at, created_at
                        FROM apply_posts
                        WHERE post_id IN(?` + strings.Repeat(",?", len(pIDs)-1) + ")"

	args := make([]interface{}, len(pIDs))
	for i, p := range pIDs {
		args[i] = p
	}
	return r.fetchApplyPosts(ctx, query, args...)
}

func (r *applyPostRepo) CountApplyPostsByPostID(ctx context.Context, postID int64) (int64, error) {
	query := `SELECT COUNT(*)
                     FROM apply_posts
										 WHERE post_id = ?`

	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var cnt int64
	rows, err := stmt.QueryContext(ctx, postID)
	if err != nil {
		return 0, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Println(err)
		}
	}()
	for rows.Next() {
		if err := rows.Scan(&cnt); err != nil {
			return 0, err
		}
	}
	return cnt, nil
}

func (r *applyPostRepo) CreateApplyPost(ctx context.Context, p *models.ApplyPost) error {
	query := `INSERT apply_posts SET post_id=?, user_id=?, updated_at=?, created_at=?`
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, p.PostID, p.UserID, p.UpdatedAt, p.CreatedAt)
	if err != nil {
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCnt != 1 {
		return fmt.Errorf("expected %d row affected, got %d rows affected", 1, rowCnt)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = lastID
	return nil
}

func (r *applyPostRepo) DeleteApplyPost(ctx context.Context, id int64) error {
	query := `DELETE FROM apply_posts WHERE id = ?`

	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

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
