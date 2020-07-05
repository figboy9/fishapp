package middleware

import (
	"time"

	"github.com/ezio1119/fishapp-image/conf"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
)

func (*middleware) UnaryLogingInterceptor() grpc.UnaryServerInterceptor {
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

func (*middleware) StreamLogingInterceptor() grpc.StreamServerInterceptor {
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

// func (*middleware) LoggerInterceptor() grpc.UnaryServerInterceptor {
// 	var logrusLogger *logrus.Logger
// 	logrusEntry := logrus.NewEntry(logrusLogger)
// 	// Shared options for the logger, with a custom duration to log field function.
// 	opts := []grpc_logrus.Option{
// 		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
// 			return "grpc.time_ns", duration.Nanoseconds()
// 		}),
// 	}

// 	return grpc_logrus.UnaryServerInterceptor(logrusEntry, opts...)
// }
