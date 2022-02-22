package models

var RefreshTokenName string = "rtokens"

type RefreshToken struct {
	Token string `db:"token"`
}
