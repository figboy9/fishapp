package sqlhandler

import (
	"context"
	"database/sql"
)

type SqlHandler interface {
	BeginTx(ctx context.Context) (context.Context, error)
	Roolback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}
