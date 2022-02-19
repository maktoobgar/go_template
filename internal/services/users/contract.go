package user_service

import "github.com/maktoobgar/go_template/internal/models"

type UserService interface {
	GetUser(username string) (*models.User, error)
	CreateUser(username string, password string, display_name string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	CheckPasswordHash(password, hash string) bool
	HashPassword(password string) string
}
