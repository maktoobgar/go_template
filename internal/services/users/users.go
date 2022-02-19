package user_service

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/errors/messages"
	"golang.org/x/crypto/bcrypt"
)

type userService struct{}

var errDataNotFound = errors.New(errors.NotFoundStatus, messages.ErrDataNotFound)
var errWrongData = errors.New(errors.InvalidStatus, messages.ErrWrongData)
var errUnexpected = errors.New(errors.UnexpectedStatus, messages.ErrUnexpected)

var instance = &userService{}

func (obj *userService) HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func (obj *userService) CheckPasswordHash(password, hash string) bool {
	return nil == bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (obj *userService) CreateUser(username string, password string, display_name string) (*models.User, error) {
	user := &models.User{
		Username:    username,
		Password:    obj.HashPassword(password),
		DisplayName: display_name,
		JoinedDate:  time.Now(),
	}

	_, err := g.DB.Insert(models.UserName).Rows([]*models.User{user}).Executor().Exec()
	if err != nil {
		return nil, errUnexpected
	}

	user, err = obj.GetUser(username)
	if err != nil {
		return nil, errUnexpected
	}

	return user, nil
}

func (obj *userService) GetUser(username string) (*models.User, error) {
	user := &models.User{}
	ok, err := g.DB.From(models.UserName).Limit(1).Where(goqu.Ex{
		"username": username,
	}).Executor().ScanStruct(user)

	if !ok || err != nil || user == nil {
		return nil, errDataNotFound
	}

	return user, nil
}

func (obj *userService) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	ok, err := g.DB.From(models.UserName).Limit(1).Where(goqu.Ex{
		"id": id,
	}).Executor().ScanStruct(user)

	if !ok || err != nil || user == nil {
		return nil, errDataNotFound
	}

	return user, nil
}

func New() UserService {
	return instance
}
