package user_service

import (
	"context"
	"database/sql"
	"time"

	"github.com/maktoobgar/go_template/internal/contract"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/internal/repositories"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

var instance = &userService{}

func (obj *userService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (obj *userService) CheckPasswordHash(password, hash string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (obj *userService) CreateUser(db *sql.DB, ctx context.Context, username string, password string, display_name string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{
		UserCore: models.UserCore{
			Username:    username,
			DisplayName: display_name,
			JoinedDate:  time.Now(),
		},
		Password: obj.HashPassword(password),
	}

	query := repositories.InsertInto(user.Name(), user, ctx)
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, translate("SignUpFailure")))
	}
	user.ID, _ = result.LastInsertId()

	return user
}

func (obj *userService) GetUser(db *sql.DB, ctx context.Context, username string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{}
	query := repositories.Select(user.Name(), map[string]any{
		"username": username,
	}, ctx)
	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		panic(errors.New(errors.NotFoundStatus, errors.Resend, translate("UserNotFound")))
	}

	return user
}

func (obj *userService) SafeGetUser(db *sql.DB, ctx context.Context, username string) *models.User {
	user := &models.User{}
	query := repositories.Select(user.Name(), map[string]any{
		"username": username,
	}, ctx)
	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		return nil
	}

	return user
}

func (obj *userService) GetUserByID(db *sql.DB, ctx context.Context, id string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{}
	query := repositories.Select(user.Name(), map[string]any{
		"id": id,
	}, ctx)
	err := db.QueryRowContext(ctx, query).Scan(user)
	if err != nil {
		panic(errors.New(errors.NotFoundStatus, errors.Resend, translate("UserNotFound")))
	}

	return user
}

func New() contract.UserService {
	return instance
}
