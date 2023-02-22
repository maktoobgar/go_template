package contract

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/models"
)

type TokenService interface {
	// Returns token string, expire time and error
	CreateAccessToken(user *models.User, ctx context.Context) (string, time.Time)
	// Returns refresh token object, expire time and error
	CreateRefreshToken(db *goqu.Database, ctx context.Context, user *models.User) (string, time.Time)
	// Returns a refresh token from database
	GetRefreshToken(db *goqu.Database, ctx context.Context, token string) *models.RefreshToken
	// Returns a refresh token from database without panic
	SafeGetRefreshToken(db *goqu.Database, ctx context.Context, token string) *models.RefreshToken
	// Deletes a refresh token from database
	DeleteRefreshToken(db *goqu.Database, ctx context.Context, token string)
}
