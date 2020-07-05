package sqlhandler

import (
	"context"
	"database/sql"
	"sync"

	"github.com/google/uuid"
)

type contextKey string

const (
	sqlTxMapCtxKey contextKey = "txMapKey"
	isTxCtxKey     contextKey = "isTx"
)

type sqlHandler struct {
	db    *sql.DB
	txMap map[string]*sql.Tx
	mu    sync.RWMutex
}

func NewSqlHandler(db *sql.DB) SqlHandler {
	return &sqlHandler{
		db:    db,
		txMap: map[string]*sql.Tx{},
		mu:    sync.RWMutex{},
	}
}

func (h *sqlHandler) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if h.isTx(ctx) {
		return h.getTx(h.getTxMapKeyFromCtx(ctx)).PrepareContext(ctx, query)
	}
	return h.db.PrepareContext(ctx, query)
}

// ctxの中にtxMapKeyをいれて、sqlhandlerのtxMapにsql.Txを入れる
func (h *sqlHandler) BeginTx(ctx context.Context) (context.Context, error) {

	tx, err := h.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	txMapKey := uuid.New().String()
	h.setTx(tx, txMapKey)

	return h.setTxMapKeyToCtx(h.setIsTxTrueToCtx(ctx), txMapKey), nil
}

func (h *sqlHandler) Roolback(ctx context.Context) (context.Context, error) {
	txMapKey := h.getTxMapKeyFromCtx(ctx)

	defer h.deleteTx(txMapKey)

	ctx = h.setIsTxfalseToCtx(ctx)

	if err := h.getTx(txMapKey).Rollback(); err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (h *sqlHandler) Commit(ctx context.Context) (context.Context, error) {
	txMapKey := h.getTxMapKeyFromCtx(ctx)
	defer h.deleteTx(txMapKey)

	ctx = h.setIsTxfalseToCtx(ctx)

	if err := h.getTx(txMapKey).Commit(); err != nil {
		return ctx, err
	}

	return ctx, nil
}

func (h *sqlHandler) setTx(tx *sql.Tx, txMapKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.txMap[txMapKey] = tx
}

func (h *sqlHandler) getTx(txMapKey string) *sql.Tx {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.txMap[txMapKey]
}

func (h *sqlHandler) deleteTx(txMapKey string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.txMap, txMapKey)
}

func (h *sqlHandler) isTx(ctx context.Context) bool {
	if isTx, ok := ctx.Value(isTxCtxKey).(bool); ok {
		return isTx
	}
	return false
}

func (h *sqlHandler) setTxMapKeyToCtx(ctx context.Context, txMapKey string) context.Context {
	return context.WithValue(ctx, sqlTxMapCtxKey, txMapKey)
}

func (h *sqlHandler) getTxMapKeyFromCtx(ctx context.Context) string {
	return ctx.Value(sqlTxMapCtxKey).(string)
}

func (h *sqlHandler) setIsTxTrueToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, isTxCtxKey, true)
}

func (h *sqlHandler) setIsTxfalseToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, isTxCtxKey, false)
}
