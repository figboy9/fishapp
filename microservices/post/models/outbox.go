package models

import "time"

type Outbox struct {
	ID            string
	EventType     string
	EventData     []byte
	AggregateID   string
	AggregateType string
	Channel       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
