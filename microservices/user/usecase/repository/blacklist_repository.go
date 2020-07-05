package repository

import "time"

type BlackListRepository interface {
	SetNX(jti string, exp time.Duration) (bool, error)
	Exists(t string) (int64, error)
}
