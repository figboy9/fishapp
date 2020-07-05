package domain

import (
	"github.com/siddontang/go-mysql/mysql"
)

type BinlogEvent struct {
	LastPos    mysql.Position
	PosSubject string
	Event      []interface{}
}
