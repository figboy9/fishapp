package models

import "time"

type ApplyPost struct {
	ID        int64
	UserID    int64
	PostID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
