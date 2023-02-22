package token_service

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

type tokenService struct{}

var instance = &tokenService{}

func (obj *tokenService) CreateAccessToken(user *models.User, ctx context.Context) (string, time.Time) {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
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
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("GenerationTokenFailed")))
	}

	return tokenString, expirationTime
}

func (obj *tokenService) CreateRefreshToken(db *goqu.Database, ctx context.Context, user *models.User) (string, time.Time) {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
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
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("GenerationTokenFailed")))
	}

	rows := []models.RefreshToken{
		{Token: tokenString},
	}

	_, err = db.Insert(models.RefreshTokenName).Rows(rows).Executor().Exec()
	if err != nil {
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("CreationRefreshTokenFailed")))
	}

	return tokenString, expirationTime
}

func (obj *tokenService) GetRefreshToken(db *goqu.Database, ctx context.Context, token string) *models.RefreshToken {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	refreshToken := &models.RefreshToken{}
	ok, err := db.From(models.RefreshTokenName).Limit(1).Where(exp.Ex{
		"token": token,
	}).Executor().ScanStruct(refreshToken)
	if !ok || err != nil {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("UsedRefreshToken")))
	}

	return refreshToken
}

func (obj *tokenService) SafeGetRefreshToken(db *goqu.Database, ctx context.Context, token string) *models.RefreshToken {
	refreshToken := &models.RefreshToken{}
	ok, err := db.From(models.RefreshTokenName).Limit(1).Where(exp.Ex{
		"token": token,
	}).Executor().ScanStruct(refreshToken)
	if !ok || err != nil {
		return nil
	}

	return refreshToken
}

func (obj *tokenService) DeleteRefreshToken(db *goqu.Database, ctx context.Context, token string) {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	_, err := db.Delete(models.RefreshTokenName).Where(exp.Ex{
		"token": token,
	}).Executor().Exec()
	if err != nil {
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("DeletionRefreshTokenFailed")))
	}
}

func New() contract.TokenService {
	return instance
}
