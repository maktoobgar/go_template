package user_service

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/contract"
	"github.com/maktoobgar/go_template/internal/models"
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

func (obj *userService) CreateUser(db *goqu.Database, ctx context.Context, username string, password string, display_name string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{
		UserCore: models.UserCore{
			Username:    username,
			DisplayName: display_name,
			JoinedDate:  time.Now(),
		},
		Password: obj.HashPassword(password),
	}

	_, err := db.Insert(models.UserName).Rows([]*models.User{user}).Executor().Exec()
	if err != nil {
		panic(errors.New(errors.InvalidStatus, errors.Resend, translate("SignUpFailure")))
	}

	user = obj.GetUser(db, ctx, username)
	return user
}

func (obj *userService) GetUser(db *goqu.Database, ctx context.Context, username string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{}
	ok, err := db.From(models.UserName).Limit(1).Where(goqu.Ex{
		"username": username,
	}).Executor().ScanStruct(user)

	if !ok || err != nil {
		panic(errors.New(errors.NotFoundStatus, errors.Resend, translate("UserNotFound")))
	}

	return user
}

func (obj *userService) SafeGetUser(db *goqu.Database, ctx context.Context, username string) *models.User {
	user := &models.User{}
	ok, err := db.From(models.UserName).Limit(1).Where(goqu.Ex{
		"username": username,
	}).Executor().ScanStruct(user)

	if !ok || err != nil {
		return nil
	}

	return user
}

func (obj *userService) GetUserByID(db *goqu.Database, ctx context.Context, id string) *models.User {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	user := &models.User{}
	ok, err := db.From(models.UserName).Limit(1).Where(goqu.Ex{
		"id": id,
	}).Executor().ScanStruct(user)

	if !ok || err != nil || user == nil {
		panic(errors.New(errors.NotFoundStatus, errors.Resend, translate("UserNotFound")))
	}

	return user
}

func New() contract.UserService {
	return instance
}
