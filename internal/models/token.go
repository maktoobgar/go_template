package models

var RefreshTokenName string = "rtokens"

type RefreshToken struct {
	Token string `db:"token"`
}

func (t *RefreshToken) Name() string {
	return RefreshTokenName
}
