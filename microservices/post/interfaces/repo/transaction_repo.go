package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/usecase/repo"
)

type transactionRepo struct {
	SqlHandler
}

func NewTransactionRepo(h SqlHandler) repo.TransactionRepo {
	return &transactionRepo{h}
}

func (r *transactionRepo) BeginTx(ctx context.Context) (context.Context, error) {
	return r.SqlHandler.BeginTx(ctx)
}

func (r *transactionRepo) Roolback(ctx context.Context) (context.Context, error) {
	return r.SqlHandler.Roolback(ctx)
}

func (r *transactionRepo) Commit(ctx context.Context) (context.Context, error) {
	return r.SqlHandler.Commit(ctx)
}
