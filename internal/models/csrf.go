package models

import (
	"time"
)

var CSRFName string = "csrfs"

type CSRF struct {
	ID         int       `db:"id" goqu:"skipinsert"`
	Key        string    `db:"key"`
	Value      string    `db:"value"`
	ExpireDate time.Time `db:"expire_date" goqu:"skipupdate"`
}
