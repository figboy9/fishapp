package middleware

import (
	"context"

	"github.com/ezio1119/fishapp-chat/conf"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (*Middleware) UnaryRecoveryInterceptor() grpc.UnaryServerInterceptor {
	if conf.C.Sv.Debug {
		return emptyUnalyIntercepter()
	}

	customFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Internal, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.

	return grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc))
}

func (*Middleware) StreamRecoveryInterceptor() grpc.StreamServerInterceptor {
	if conf.C.Sv.Debug {
		return emptyStreamIntercepter()
	}

	customFunc := func(p interface{}) (err error) {
		return status.Errorf(codes.Internal, "panic triggered: %v", p)
	}
	// Shared options for the logger, with a custom gRPC code to log level function.

	return grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(customFunc))
}

func emptyStreamIntercepter() grpc.StreamServerInterceptor {
	return grpc.StreamServerInterceptor(func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		return handler(srv, ss)
	})
}

func emptyUnalyIntercepter() grpc.UnaryServerInterceptor {
	return grpc.UnaryServerInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		return handler(ctx, req)
	})
}
