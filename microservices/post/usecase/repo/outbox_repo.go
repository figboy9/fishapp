package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/models"
)

type OutboxRepo interface {
	CreateOutbox(ctx context.Context, o *models.Outbox) error
}
