package repo

import "context"

type TransactionRepo interface {
	BeginTx(ctx context.Context) (context.Context, error)
	Roolback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
}
