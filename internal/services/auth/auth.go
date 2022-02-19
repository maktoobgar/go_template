package auth_service

import (
	"github.com/maktoobgar/go_template/internal/models"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/errors/messages"
)

type authService struct{}

var errDataNotFound = errors.New(errors.NotFoundStatus, messages.ErrDataNotFound)
var errWrongData = errors.New(errors.InvalidStatus, messages.ErrWrongData)
var errUnexpected = errors.New(errors.UnexpectedStatus, messages.ErrUnexpected)

var instance = &authService{}

func (obj *authService) authenticate(username string, password string) (*models.User, error) {
	uService := user_service.New()
	user, err := uService.GetUser(username)
	if err != nil {
		return nil, err
	}

	if !uService.CheckPasswordHash(password, user.Password) {
		return nil, errWrongData
	}

	return user, nil
}

func (obj *authService) SignIn(username string, password string) (*models.User, error) {
	user, err := obj.authenticate(username, password)
	if err != nil || user == nil {
		return nil, err
	}

	return user, nil
}

func (obj *authService) SignUp(username string, password string, display_name string) (*models.User, error) {
	uService := user_service.New()
	_, err := uService.GetUser(username)
	if err == nil {
		return nil, errDataNotFound
	}

	user, err := uService.CreateUser(username, password, display_name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func New() AuthService {
	return instance
}
