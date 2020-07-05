package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/models"
)

type ApplyPostRepo interface {
	GetApplyPostByID(ctx context.Context, id int64) (*models.ApplyPost, error)
	ListApplyPostsByUserID(ctx context.Context, userID int64) ([]*models.ApplyPost, error)
	ListApplyPostsByPostID(ctx context.Context, postID int64) ([]*models.ApplyPost, error)
	BatchGetApplyPostsByPostIDs(ctx context.Context, postIDs []int64) ([]*models.ApplyPost, error)
	CountApplyPostsByPostID(ctx context.Context, postID int64) (int64, error)
	CreateApplyPost(ctx context.Context, p *models.ApplyPost) error
	DeleteApplyPost(ctx context.Context, id int64) error
}
