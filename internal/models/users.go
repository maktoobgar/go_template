package models

import "time"

var UserName string = "users"

type UserCore struct {
	ID          int64     `db:"id" goqu:"skipinsert"`
	Username    string    `db:"username"`
	DisplayName string    `db:"display_name"`
	FirstName   string    `db:"first_name"`
	LastName    string    `db:"last_name"`
	JoinedDate  time.Time `db:"joined_date" goqu:"skipupdate"`
}

type User struct {
	UserCore
	Password string `db:"password"`
}

// Returns name of the table in database
func (u *User) Name() string {
	return UserName
}
