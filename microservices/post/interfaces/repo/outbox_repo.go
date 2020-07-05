package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/usecase/repo"
)

type outboxRepo struct {
	SqlHandler
}

func NewOutboxRepo(s SqlHandler) repo.OutboxRepo {
	return &outboxRepo{s}
}

func (r *outboxRepo) CreateOutbox(ctx context.Context, o *models.Outbox) error {
	query := `INSERT outbox SET id=?, event_type=?, event_data=?, aggregate_id=?, aggregate_type=?, channel=?, updated_at=?, created_at=?`

	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, o.ID, o.EventType, o.EventData, o.AggregateID, o.AggregateType, o.Channel, o.UpdatedAt, o.CreatedAt)
	if err != nil {
		return err
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if int(rowCnt) != 1 {
		return err
	}

	return nil
}
