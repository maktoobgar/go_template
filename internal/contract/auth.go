package contract

import (
	"context"
	"database/sql"

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
	SignIn(db *sql.DB, ctx context.Context, username string, password string) *models.User
	// Signs up a user with minimum required fields
	SignUp(db *sql.DB, ctx context.Context, username string, password string, display_name string) *models.User
}
