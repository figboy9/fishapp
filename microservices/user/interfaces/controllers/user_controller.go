package controllers

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/ezio1119/fishapp-user/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
)

type userController struct {
	userInteractor interactor.UserUsecase
	authFunc       func(ctx context.Context, tokenType domain.TokenType) (context.Context, error) // grpc.StreamServerInterceptorでcontextを渡せないためコントローラー層で認証する
}

func NewUserController(
	a interactor.UserUsecase,
	authFunc func(ctx context.Context, tokenType domain.TokenType) (context.Context, error),
) pb.UserServiceServer {
	return &userController{a, authFunc}
}

func (c *userController) CurrentUser(ctx context.Context, in *pb.CurrentUserReq) (*pb.User, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	u, err := c.userInteractor.GetUser(ctx, uID)
	if err != nil {
		return nil, err
	}

	return convUserProto(u)
}

func (c *userController) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	u, err := c.userInteractor.GetUser(ctx, in.Id)
	if err != nil {

		return nil, err
	}

	return convUserProto(u)
}

func (c *userController) CreateUser(stream pb.UserService_CreateUserServer) error {
	ctx := stream.Context()

	u := &domain.User{}
	imageBuf := &bytes.Buffer{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.CreateUserReq_Info:
			sex, err := convSex(x.Info.Sex)
			if err != nil {
				return err
			}

			u.Email = x.Info.Email
			u.Password = x.Info.Password
			u.Name = x.Info.Name
			u.Introduction = x.Info.Introduction
			u.Sex = sex
		case *pb.CreateUserReq_ImageChunk:
			if _, err := imageBuf.Write(x.ImageChunk); err != nil {
				return err
			}
		}
	}

	t, err := c.userInteractor.CreateUser(ctx, u, imageBuf)
	if err != nil {
		return err
	}

	pbUser, err := convUserProto(u)
	if err != nil {
		return err
	}

	return stream.SendAndClose(&pb.CreateUserRes{
		User:      pbUser,
		TokenPair: convTokenPairProto(t),
	})
}

func (c *userController) UpdateUser(stream pb.UserService_UpdateUserServer) error {
	ctx := stream.Context()
	ctx, err := c.authFunc(ctx, domain.IdToken)
	if err != nil {
		return err
	}

	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return err
	}

	u := &domain.User{ID: uID}
	imageBuf := &bytes.Buffer{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		switch x := req.Data.(type) {
		case *pb.UpdateUserReq_Info:
			u.Email = x.Info.Email
			u.Name = x.Info.Name
			u.Introduction = x.Info.Introduction
		case *pb.UpdateUserReq_ImageChunk:
			if _, err := imageBuf.Write(x.ImageChunk); err != nil {
				return err
			}
		}
	}

	// トークンからユーザーIDを取り出しているので認可なし
	if err := c.userInteractor.UpdateUser(ctx, u, imageBuf); err != nil {
		return err
	}

	pbUser, err := convUserProto(u)
	if err != nil {
		return err
	}

	return stream.SendAndClose(pbUser)
}

func (c *userController) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordReq) (*empty.Empty, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := c.userInteractor.UpdatePassword(ctx, uID, in.OldPassword, in.NewPassword); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *userController) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	u, t, err := c.userInteractor.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	uProto, err := convUserProto(u)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{User: uProto, TokenPair: convTokenPairProto(t)}, nil
}

func (c *userController) Logout(ctx context.Context, in *pb.LogoutReq) (*empty.Empty, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.userInteractor.Logout(ctx, &claims); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *userController) RefreshIDToken(ctx context.Context, in *pb.RefreshIDTokenReq) (*pb.RefreshIDTokenRes, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	t, err := c.userInteractor.RefreshIDToken(ctx, &claims)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshIDTokenRes{TokenPair: convTokenPairProto(t)}, nil
}
