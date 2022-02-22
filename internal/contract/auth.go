package contract

import "github.com/dgrijalva/jwt-go"

var RefreshTokenType = "1"
var AccessTokenType = "2"

type Claims struct {
	Username string `json:"username"`
	Type     string `json:"type"`
	jwt.StandardClaims
}
