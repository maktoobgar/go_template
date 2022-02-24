package contract

import "github.com/maktoobgar/go_template/internal/models"

type UserService interface {
	// Returns user model object based on username
	GetUser(username string) (*models.User, error)
	// Creates a user based on required data
	CreateUser(username string, password string, display_name string) (*models.User, error)
	// Returns user model object based on id
	GetUserByID(id string) (*models.User, error)
	// Checks if the passwords match
	CheckPasswordHash(password, hash string) bool
	// Hashes received password
	HashPassword(password string) string
}
