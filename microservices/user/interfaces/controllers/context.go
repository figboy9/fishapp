package controllers

import (
	"context"

	"github.com/ezio1119/fishapp-user/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getJwtClaimsCtx(ctx context.Context) (domain.JwtClaims, error) {
	v := ctx.Value(domain.JwtCtxKey)
	c, ok := v.(domain.JwtClaims)
	if !ok {
		return domain.JwtClaims{}, status.Errorf(codes.Internal, "failed to get jwt claims from context")
	}

	return c, nil
}
