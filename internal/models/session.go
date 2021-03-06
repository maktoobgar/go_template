package models

import (
	"time"
)

var SessionName string = "sessions"

type Session struct {
	ID         int       `db:"id" goqu:"skipinsert"`
	Key        string    `db:"key"`
	Value      string    `db:"value"`
	ExpireDate time.Time `db:"expire_date" goqu:"skipupdate"`
}
