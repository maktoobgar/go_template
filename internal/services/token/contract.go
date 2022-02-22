package token_service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/maktoobgar/go_template/internal/models"
)

type TokenService interface {
	// Returns token string, expire time and error
	CreateAccessToken(user *models.User) (string, time.Time, error)
	// Returns refresh token object, expire time and error
	CreateRefreshToken(user *models.User) (string, time.Time, error)
	// Returns a refresh token from database
	GetRefreshToken(token string) (*models.RefreshToken, error)
	// Deletes a refresh token from database
	DeleteRefreshToken(token string) error
	// Returns the key from a token object
	SignedString(token *jwt.Token) (string, error)
}
