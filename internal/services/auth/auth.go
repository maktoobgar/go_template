package auth_service

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type authService struct{}

var instance = &authService{}

func (obj *authService) authenticate(db *goqu.Database, username string, password string) (*models.User, error) {
	uService := user_service.New()
	user, err := uService.GetUser(db, username)
	if err != nil {
		return nil, err
	}

	if !uService.CheckPasswordHash(password, user.Password) {
		return nil, errors.New(errors.UnauthorizedStatus, errors.Resend, g.Trans().TranslateEN("UnMatchPassword"))
	}

	return user, nil
}

func (obj *authService) SignIn(db *goqu.Database, username string, password string) (*models.User, error) {
	user, err := obj.authenticate(db, username, password)
	if err != nil || user == nil {
		return nil, err
	}

	return user, nil
}

func (obj *authService) SignUp(db *goqu.Database, username string, password string, display_name string) (*models.User, error) {
	uService := user_service.New()
	_, err := uService.GetUser(db, username)
	if err == nil {
		return nil, errors.New(errors.InvalidStatus, errors.ReSingIn, g.Trans().TranslateEN("DuplicateUser"))
	}

	user, err := uService.CreateUser(db, username, password, display_name)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func New() contract.AuthService {
	return instance
}
