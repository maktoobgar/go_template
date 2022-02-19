package auth_service

import "github.com/maktoobgar/go_template/internal/models"

type AuthService interface {
	// Sings in a user with authenticating username and password
	SignIn(username string, password string) (*models.User, error)
	// Signs up a user with minimum required fields
	SignUp(username string, password string, display_name string) (*models.User, error)
}
