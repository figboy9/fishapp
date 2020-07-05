package interactor

import (
	"bytes"
	"context"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/usecase/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userInteractor struct {
	userRepository      repository.UserRepository
	blackListRepository repository.BlackListRepository
	imageRepository     repository.ImageRepository
	ctxTimeout          time.Duration
}

func NewUserInteractor(
	u repository.UserRepository,
	b repository.BlackListRepository,
	i repository.ImageRepository,
	t time.Duration,
) *userInteractor {
	return &userInteractor{u, b, i, t}
}

type UserUsecase interface {
	CreateUser(ctx context.Context, u *domain.User, imageBuf *bytes.Buffer) (*domain.TokenPair, error)
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, u *domain.User, imageBuf *bytes.Buffer) error
	UpdatePassword(ctx context.Context, id int64, oldPassword string, newPassword string) error
	Login(ctx context.Context, email string, pass string) (*domain.User, *domain.TokenPair, error)
	Logout(ctx context.Context, jwtClaims *domain.JwtClaims) error
	RefreshIDToken(ctx context.Context, jwtClaims *domain.JwtClaims) (*domain.TokenPair, error)
}

func (i *userInteractor) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	u, err := i.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	// getuserは認証なしのため、emailは参照できない
	u.Email = ""
	return u, nil
}

func (i *userInteractor) CreateUser(ctx context.Context, u *domain.User, imageBuf *bytes.Buffer) (*domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	encryptedPass, err := genEncryptedPass(u.Password)
	if err != nil {
		return nil, err
	}

	u.EncryptedPassword = encryptedPass

	if err := i.userRepository.CreateUser(ctx, u); err != nil {
		return nil, err
	}

	if imageBuf.Len() != 0 {
		if err := i.imageRepository.BatchCreateImages(context.Background(), u.ID, []*bytes.Buffer{imageBuf}); err != nil {
			if err := i.userRepository.DeleteUser(context.Background(), u.ID); err != nil {
				return nil, err
			}
			return nil, err
		}
	}

	return genTokenPair(strconv.FormatInt(u.ID, 10))
}

func (i *userInteractor) UpdateUser(ctx context.Context, u *domain.User, imageBuf *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	oldU, err := i.userRepository.GetUser(ctx, u.ID)
	if err != nil {
		return err
	}

	if err := i.userRepository.UpdateUser(ctx, u); err != nil {
		return err
	}

	if imageBuf.Len() != 0 {
		if err := i.imageRepository.DeleteImagesByUserID(ctx, u.ID); err != nil {
			if err := i.userRepository.UpdateUser(ctx, oldU); err != nil {
				return err
			}
		}

		if err := i.imageRepository.BatchCreateImages(ctx, u.ID, []*bytes.Buffer{imageBuf}); err != nil {
			if err := i.userRepository.UpdateUser(ctx, oldU); err != nil {
				return err
			}
		}
	}

	return nil
}

func (i *userInteractor) UpdatePassword(ctx context.Context, id int64, oldPwd string, newPwd string) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	res, err := i.userRepository.GetUser(ctx, id)
	if err != nil {
		return err
	}

	if err := compareHashAndPass(res.EncryptedPassword, oldPwd); err != nil {
		return err
	}

	encryptedPass, err := genEncryptedPass(newPwd)
	if err != nil {
		return status.Error(codes.Unauthenticated, err.Error())
	}

	u := &domain.User{
		ID:                id,
		EncryptedPassword: encryptedPass,
	}

	return i.userRepository.UpdateUser(ctx, u)
}

func (i *userInteractor) Login(ctx context.Context, email string, pass string) (*domain.User, *domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	u, err := i.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}

	if err := compareHashAndPass(u.EncryptedPassword, pass); err != nil {
		return nil, nil, status.Errorf(codes.Unauthenticated, err.Error())
	}

	tokenPair, err := genTokenPair(strconv.FormatInt(u.ID, 10))
	if err != nil {
		return nil, nil, err
	}

	return u, tokenPair, nil
}

func (i *userInteractor) Logout(ctx context.Context, c *domain.JwtClaims) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	exp := time.Unix(c.ExpiresAt, 0).Sub(time.Now())
	if _, err := i.blackListRepository.SetNX(c.Id, exp); err != nil {
		return err
	}

	return nil
}

func (i *userInteractor) RefreshIDToken(ctx context.Context, c *domain.JwtClaims) (*domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	res, err := i.blackListRepository.Exists(c.Id)
	if err != nil {
		return nil, err
	}
	if res == 1 {
		return nil, status.Error(codes.Unauthenticated, "token is blacklisted")
	}

	tokenPair, err := genTokenPair(c.User.ID)
	if err != nil {
		return nil, err
	}
	exp := time.Unix(c.ExpiresAt, 0).Sub(time.Now()) // トークンの有効期限 - 現在の時間
	if _, err := i.blackListRepository.SetNX(c.Id, exp); err != nil {
		return nil, err
	}

	return tokenPair, nil
}
