package contract

import (
	"context"
	"database/sql"
	"time"

	"github.com/maktoobgar/go_template/internal/models"
)

type TokenService interface {
	// Returns token string, expire time and error
	CreateAccessToken(user *models.User, ctx context.Context) (string, time.Time)
	// Returns refresh token object, expire time and error
	CreateRefreshToken(db *sql.DB, ctx context.Context, user *models.User) (string, time.Time)
	// Returns a refresh token from database
	GetRefreshToken(db *sql.DB, ctx context.Context, token string) *models.RefreshToken
	// Returns a refresh token from database without panic
	SafeGetRefreshToken(db *sql.DB, ctx context.Context, token string) *models.RefreshToken
	// Deletes a refresh token from database
	DeleteRefreshToken(db *sql.DB, ctx context.Context, token string)
}
