package repository

import (
	"context"

	"github.com/ezio1119/fishapp-user/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *domain.User) error
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, u *domain.User) error
	DeleteUser(ctx context.Context, id int64) error
}
