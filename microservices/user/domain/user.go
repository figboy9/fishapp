package domain

import "time"

type User struct {
	ID                int64
	Email             string
	Password          string `gorm:"-"`
	Name              string
	Introduction      string
	Sex               Sex
	EncryptedPassword string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Sex int64

const (
	Male Sex = iota + 1
	Female
)
