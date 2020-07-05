package middleware

import (
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"google.golang.org/grpc"
)

func (*Middleware) UnaryValidationInterceptor() grpc.UnaryServerInterceptor {
	return grpc_validator.UnaryServerInterceptor()
}

func (*Middleware) StreamValidationInterceptor() grpc.StreamServerInterceptor {
	return grpc_validator.StreamServerInterceptor()
}
