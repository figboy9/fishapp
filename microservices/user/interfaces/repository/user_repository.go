package repository

import (
	"context"
	"fmt"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userRepository struct {
	conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *userRepository {
	return &userRepository{conn}
}

func (r *userRepository) CreateUser(ctx context.Context, u *domain.User) error {
	result := r.conn.Create(u)
	if err := result.Error; err != nil {
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return status.Errorf(codes.Internal, "%d rows affected", rows)
	}
	return nil
}

func (r *userRepository) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	var u domain.User
	if err := r.conn.Take(&u, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "user with id='%d' is not found", id)
		}
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User
	if err := r.conn.Where("email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "user with email='%s' is not found", email)
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, u *domain.User) error {
	result := r.conn.Model(u).Updates(u) // SET 'id'も含まれてしまう
	if err := result.Error; err != nil {
		fmt.Printf("error: %#v\n", err)
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}

	return r.conn.Take(u).Error
}

func (r *userRepository) DeleteUser(ctx context.Context, id int64) error {
	return r.conn.Delete(&domain.User{ID: id}).Error
}
