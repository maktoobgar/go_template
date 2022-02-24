package contract

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/maktoobgar/go_template/internal/models"
)

var RefreshTokenType = "1"
var AccessTokenType = "2"

type Claims struct {
	Username string `json:"username"`
	Type     string `json:"type"`
	jwt.StandardClaims
}

type AuthService interface {
	// Sings in a user with authenticating username and password
	SignIn(username string, password string) (*models.User, error)
	// Signs up a user with minimum required fields
	SignUp(username string, password string, display_name string) (*models.User, error)
}
