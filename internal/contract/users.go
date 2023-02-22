package contract

import (
	"context"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/models"
)

type UserService interface {
	// Returns user model object based on username
	GetUser(db *goqu.Database, ctx context.Context, username string) *models.User
	// Returns user model object based on username safe
	SafeGetUser(db *goqu.Database, ctx context.Context, username string) *models.User
	// Creates a user based on required data
	CreateUser(db *goqu.Database, ctx context.Context, username string, password string, display_name string) *models.User
	// Returns user model object based on id
	GetUserByID(db *goqu.Database, ctx context.Context, id string) *models.User
	// Checks if the passwords match
	CheckPasswordHash(password, hash string) bool
	// Hashes received password
	HashPassword(password string) string
}
