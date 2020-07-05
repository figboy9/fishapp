package domain

import (
	"time"
)

type Message struct {
	ID        int64
	Body      string
	RoomID    int64
	UserID    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// func (m *Message) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(m)
// }

// func (m *Message) UnmarshalBinary(data []byte) error {
// 	return json.Unmarshal(data, m)
// }
