package middleware

import (
	"google.golang.org/grpc"
)

type Middleware interface {
	UnaryLogingInterceptor() grpc.UnaryServerInterceptor
	UnaryRecoveryInterceptor() grpc.UnaryServerInterceptor
	UnaryValidationInterceptor() grpc.UnaryServerInterceptor

	StreamLogingInterceptor() grpc.StreamServerInterceptor
	StreamRecoveryInterceptor() grpc.StreamServerInterceptor
	StreamValidationInterceptor() grpc.StreamServerInterceptor
}

type middleware struct{}

func InitMiddleware() Middleware {
	return &middleware{}
}
