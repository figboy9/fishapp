package domain

import (
	"time"
)

type Room struct {
	ID        int64
	PostID    int64
	Messages  []*Message
	Members   []*Member
	CreatedAt time.Time
	UpdatedAt time.Time
}
