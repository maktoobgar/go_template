package user_service

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
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

func (obj *userService) CreateUser(db *goqu.Database, username string, password string, display_name string) (*models.User, error) {
	user := &models.User{
		Username:    username,
		Password:    obj.HashPassword(password),
		DisplayName: display_name,
		JoinedDate:  time.Now(),
	}

	_, err := db.Insert(models.UserName).Rows([]*models.User{user}).Executor().Exec()
	if err != nil {
		return nil, errors.New(errors.InvalidStatus, errors.Resend, g.Translator.TranslateEN("SignUpFailure"))
	}

	user, err = obj.GetUser(db, username)
	if err != nil {
		return nil, errors.New(errors.InvalidStatus, errors.Resend, g.Translator.TranslateEN("SignUpFailure"))
	}

	return user, nil
}

func (obj *userService) GetUser(db *goqu.Database, username string) (*models.User, error) {
	user := &models.User{}
	ok, err := db.From(models.UserName).Limit(1).Where(goqu.Ex{
		"username": username,
	}).Executor().ScanStruct(user)

	if !ok || err != nil {
		return nil, errors.New(errors.NotFoundStatus, errors.Resend, g.Translator.TranslateEN("UserNotFound"))
	}

	return user, nil
}

func (obj *userService) GetUserByID(db *goqu.Database, id string) (*models.User, error) {
	user := &models.User{}
	ok, err := db.From(models.UserName).Limit(1).Where(goqu.Ex{
		"id": id,
	}).Executor().ScanStruct(user)

	if !ok || err != nil || user == nil {
		return nil, errors.New(errors.NotFoundStatus, errors.Resend, g.Translator.TranslateEN("UserNotFound"))
	}

	return user, nil
}

func New() contract.UserService {
	return instance
}
