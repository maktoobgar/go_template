package token_service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type tokenService struct{}

var instance = &tokenService{}

func (obj *tokenService) CreateAccessToken(user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &contract.Claims{
		Username: user.Username,
		Type:     contract.AccessTokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.SecretKey)
	if err != nil {
		return "", expirationTime, errors.New(errors.UnexpectedStatus, errors.ReSingIn, g.Translator.TranslateEN("GenerationTokenFailed"))
	}

	return tokenString, expirationTime, nil
}

func (obj *tokenService) CreateRefreshToken(db *goqu.Database, user *models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add((24 * time.Hour) * 7)

	claims := &contract.Claims{
		Username: user.Username,
		Type:     contract.RefreshTokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(g.SecretKey)
	if err != nil {
		return "", expirationTime, errors.New(errors.UnexpectedStatus, errors.ReSingIn, g.Translator.TranslateEN("GenerationTokenFailed"))
	}

	rows := []models.RefreshToken{
		{Token: tokenString},
	}

	_, err = db.Insert(models.RefreshTokenName).Rows(rows).Executor().Exec()
	if err != nil {
		return "", expirationTime, errors.New(errors.UnexpectedStatus, errors.ReSingIn, g.Translator.TranslateEN("CreationRefreshTokenFailed"))
	}

	return tokenString, expirationTime, nil
}

func (obj *tokenService) GetRefreshToken(db *goqu.Database, token string) (*models.RefreshToken, error) {
	refreshToken := &models.RefreshToken{}
	ok, err := db.From(models.RefreshTokenName).Limit(1).Where(exp.Ex{
		"token": token,
	}).Executor().ScanStruct(refreshToken)
	if !ok || err != nil {
		return nil, errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("UsedRefreshToken"))
	}

	return refreshToken, nil
}

func (obj *tokenService) DeleteRefreshToken(db *goqu.Database, token string) error {
	_, err := db.Delete(models.RefreshTokenName).Where(exp.Ex{
		"token": token,
	}).Executor().Exec()
	if err != nil {
		return errors.New(errors.UnexpectedStatus, errors.ReSingIn, g.Translator.TranslateEN("DeletionRefreshTokenFailed"))
	}
	return nil
}

func New() contract.TokenService {
	return instance
}
