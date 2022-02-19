package models

import "time"

var UserName string = "users"

type User struct {
	ID          int       `db:"id" goqu:"skipinsert"`
	Username    string    `db:"username"`
	Password    string    `db:"password"`
	DisplayName string    `db:"display_name"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	Age         int       `db:"age"`
	JoinedDate  time.Time `db:"joined_date" goqu:"skipupdate"`
}

type NotUser struct {
	ID       int    `db:"id"`
	Password string `db:"password"`
}

func (obj *User) Clean(except ...string) map[string]interface{} {
	return clean(obj, &NotUser{}, except...)
}
