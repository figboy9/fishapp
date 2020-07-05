package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ezio1119/fishapp-post/models"
	"github.com/ezio1119/fishapp-post/usecase/repo"
)

type sagaInstanceRepo struct {
	SqlHandler
}

func NewSagaInstanceRepo(h SqlHandler) repo.SagaInstanceRepo {
	return &sagaInstanceRepo{h}
}

func (r *sagaInstanceRepo) CreateSagaInstance(ctx context.Context, i *models.SagaInstance) error {
	query := `INSERT saga_instance SET id=?, saga_type=?, saga_data=?, current_state=?, updated_at=?, created_at=?`
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, i.ID, i.SagaType, i.SagaData, i.CurrentState, i.UpdatedAt, i.CreatedAt)
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

func (r *sagaInstanceRepo) UpdateSagaInstance(ctx context.Context, i *models.SagaInstance) error {
	query := `UPDATE saga_instance SET saga_data=?, current_state=?, updated_at=? WHERE id=?`
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, i.SagaData, i.CurrentState, i.UpdatedAt, i.ID)
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

func (r *sagaInstanceRepo) GetSagaInstance(ctx context.Context, sagaID string) (*models.SagaInstance, error) {
	query := `SELECT id, saga_type, saga_data, current_state, updated_at, created_at FROM saga_instance WHERE id=?`
	stmt, err := r.SqlHandler.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	i := &models.SagaInstance{}

	err = stmt.QueryRowContext(ctx, sagaID).Scan(&i.ID, &i.SagaType, &i.SagaData, &i.CurrentState, &i.UpdatedAt, &i.CreatedAt)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("no saga_instance with id %s", sagaID)
	case err != nil:
		return nil, err
	}

	return i, nil
}
