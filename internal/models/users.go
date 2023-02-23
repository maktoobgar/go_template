package models

import "time"

var UserName string = "users"

type UserCore struct {
	ID          int64     `db:"id" skipInsert:"+" json:"id"`
	Username    string    `db:"username" json:"username"`
	DisplayName string    `db:"display_name" json:"display_name"`
	FirstName   string    `db:"first_name" json:"first_name"`
	LastName    string    `db:"last_name" json:"last_name"`
	JoinedDate  time.Time `db:"joined_date" skipUpdate:"+" json:"joined_date"`
}

type User struct {
	UserCore `db:"-"`
	Password string `db:"password" json:"-"`
}

// Returns name of the table in database
func (u *User) Name() string {
	return UserName
}
