package middleware

import (
	"time"

	"github.com/ezio1119/fishapp-chat/conf"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func (*Middleware) UnaryLogingInterceptor() grpc.UnaryServerInterceptor {
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	var zapLogger *zap.Logger
	if conf.C.Sv.Debug {
		zapLogger, _ = zap.NewDevelopment()
	} else {
		zapLogger, _ = zap.NewProduction()
	}

	grpc_zap.ReplaceGrpcLogger(zapLogger)
	return grpc_zap.UnaryServerInterceptor(zapLogger, opts...)
}

func (*Middleware) StreamLogingInterceptor() grpc.StreamServerInterceptor {
	opts := []grpc_zap.Option{
		grpc_zap.WithDurationField(func(duration time.Duration) zapcore.Field {
			return zap.Int64("grpc.time_ns", duration.Nanoseconds())
		}),
	}
	var zapLogger *zap.Logger
	if conf.C.Sv.Debug {
		zapLogger, _ = zap.NewDevelopment()
	} else {
		zapLogger, _ = zap.NewProduction()
	}

	grpc_zap.ReplaceGrpcLogger(zapLogger)
	return grpc_zap.StreamServerInterceptor(zapLogger, opts...)
}
