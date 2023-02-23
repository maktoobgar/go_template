package auth_service

import (
	"context"
	"database/sql"

	"github.com/maktoobgar/go_template/internal/contract"
	"github.com/maktoobgar/go_template/internal/models"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

type authService struct{}

var instance = &authService{}

func (obj *authService) authenticate(db *sql.DB, ctx context.Context, username string, password string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	uService := user_service.New()
	user := uService.GetUser(db, ctx, username)

	if !uService.CheckPasswordHash(password, user.Password) {
		panic(errors.New(errors.UnauthorizedStatus, errors.Resend, translate("UnMatchPassword")))
	}

	return user
}

func (obj *authService) SignIn(db *sql.DB, ctx context.Context, username string, password string) *models.User {
	return obj.authenticate(db, ctx, username, password)
}

func (obj *authService) SignUp(db *sql.DB, ctx context.Context, username string, password string, display_name string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	uService := user_service.New()
	user := uService.SafeGetUser(db, ctx, username)
	if user == nil {
		user = uService.CreateUser(db, ctx, username, password, display_name)
	} else {
		panic(errors.New(errors.UnauthorizedStatus, errors.Resend, translate("DuplicateUser")))
	}

	return user
}

func New() contract.AuthService {
	return instance
}
