package contract

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/models"
)

type TokenService interface {
	// Returns token string, expire time and error
	CreateAccessToken(user *models.User) (string, time.Time, error)
	// Returns refresh token object, expire time and error
	CreateRefreshToken(db *goqu.Database, user *models.User) (string, time.Time, error)
	// Returns a refresh token from database
	GetRefreshToken(db *goqu.Database, token string) (*models.RefreshToken, error)
	// Deletes a refresh token from database
	DeleteRefreshToken(db *goqu.Database, token string) error
}
