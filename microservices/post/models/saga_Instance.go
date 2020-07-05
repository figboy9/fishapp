package models

import "time"

type SagaInstance struct {
	ID           string
	SagaType     string
	SagaData     []byte
	CurrentState string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
