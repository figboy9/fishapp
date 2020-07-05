package repo

import (
	"context"

	"github.com/ezio1119/fishapp-post/models"
)

type SagaInstanceRepo interface {
	GetSagaInstance(ctx context.Context, sagaID string) (*models.SagaInstance, error)
	CreateSagaInstance(ctx context.Context, i *models.SagaInstance) error
	UpdateSagaInstance(ctx context.Context, s *models.SagaInstance) error
}
