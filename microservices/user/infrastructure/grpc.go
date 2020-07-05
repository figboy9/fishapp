package infrastructure

import (
	"context"

	"github.com/ezio1119/fishapp-user/infrastructure/middleware"
	"github.com/ezio1119/fishapp-user/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func NewGrpcServer(middLe *middleware.Middleware, uc pb.UserServiceServer) *grpc.Server {

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middLe.UnaryLogingInterceptor(),
			middLe.UnaryValidationInterceptor(),
			middLe.UnaryAuthInterceptor(),
			middLe.UnaryRecoveryInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			middLe.StreamLogingInterceptor(),
			middLe.StreamValidationInterceptor(),
			middLe.StreamRecoveryInterceptor(),
		),
	)

	pb.RegisterUserServiceServer(server, uc)
	grpc_health_v1.RegisterHealthServer(server, &healthHandler{})
	reflection.Register(server)
	return server
}

type healthHandler struct{}

func (*healthHandler) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (*healthHandler) Watch(in *grpc_health_v1.HealthCheckRequest, s grpc_health_v1.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "watch is not implemented.")
}
