package models

import (
	"bytes"
	"time"
)

type Image struct {
	ID        int64
	Name      string
	OwnerID   int64
	OwnerType OwnerType
	Buf       *bytes.Buffer
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OwnerType int64

const (
	POST OwnerType = iota + 1
	USER
	MESSAGE
)
