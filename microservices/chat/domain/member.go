package domain

import "time"

type Member struct {
	ID        int64
	RoomID    int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
