package main

import (
	"context"
	"net"
	"time"

	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/infrastructure"
	"github.com/ezio1119/fishapp-user/infrastructure/middleware"
	"github.com/ezio1119/fishapp-user/interfaces/controllers"
	"github.com/ezio1119/fishapp-user/interfaces/repository"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/ezio1119/fishapp-user/usecase/interactor"
	"github.com/go-redis/redis/v7"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	dbConn, err := infrastructure.NewGormConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	var redisC *redis.Client
	if conf.C.Sv.Debug {
		redisC, err = infrastructure.NewRedisClient()
	} else {
		redisC, err = infrastructure.NewRedisFailoverClient()
	}

	if err != nil {
		panic(err)
	}

	defer redisC.Close()

	grpcConn, err := grpc.DialContext(ctx, conf.C.API.ImageURL, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer grpcConn.Close()
	imageC := pb.NewImageServiceClient(grpcConn)

	userController := controllers.NewUserController(
		interactor.NewUserInteractor(
			repository.NewUserRepository(dbConn),
			repository.NewBlackListRepository(redisC),
			repository.NewImageRepository(imageC),
			time.Duration(conf.C.Sv.Timeout)*time.Second,
		),
		middleware.AuthFunc, // grpc.StreamServerInterceptorでcontextを渡せないためコントローラー層で認証する
	)

	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), userController)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}
	if err := server.Serve(list); err != nil {
		panic(err)
	}

}
