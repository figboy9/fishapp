package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/models"
)

type PostRepo interface {
	GetPostByID(ctx context.Context, id int64) (*models.Post, error)
	ListPosts(ctx context.Context, p *models.Post, num int64, cursor int64, filter *models.PostFilter) ([]*models.Post, error)
	UpdatePost(ctx context.Context, p *models.Post) error
	CreatePost(ctx context.Context, p *models.Post) error
	DeletePost(ctx context.Context, id int64) error
}
