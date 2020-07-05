package main

import (
	"context"
	"net"

	"github.com/ezio1119/fishapp-image/conf"
	"github.com/ezio1119/fishapp-image/infrastructure"
	"github.com/ezio1119/fishapp-image/infrastructure/middleware"
	"github.com/ezio1119/fishapp-image/interfaces/controllers"
	"github.com/ezio1119/fishapp-image/interfaces/repo"
	"github.com/ezio1119/fishapp-image/usecase/interactor"
	repoI "github.com/ezio1119/fishapp-image/usecase/repo"
)

func main() {
	ctx := context.Background()

	dbConn, err := infrastructure.NewGormConn()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	var imageUploaderRepo repoI.ImageUploaderRepo

	if conf.C.Sv.Debug {
		imageUploaderRepo = repo.NewImageUploaderDevRepo()
	} else {

		gcsClient, err := infrastructure.NewGCSClient(ctx)
		if err != nil {
			panic(err)
		}
		defer gcsClient.Close()

		imageUploaderRepo = repo.NewImageUploaderRepo(gcsClient)
	}

	imageC := controllers.NewImageController(
		interactor.NewImageInteractor(
			dbConn,
			imageUploaderRepo,
		),
	)

	server := infrastructure.NewGrpcServer(middleware.InitMiddleware(), imageC)

	list, err := net.Listen("tcp", ":"+conf.C.Sv.Port)
	if err != nil {
		panic(err)
	}

	if err := server.Serve(list); err != nil {
		panic(err)
	}
}
