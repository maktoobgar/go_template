package token_service

import (
	"context"
	"database/sql"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/models"
	"github.com/maktoobgar/go_template/internal/repositories"
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

func (obj *tokenService) CreateRefreshToken(db *sql.DB, ctx context.Context, user *models.User) (string, time.Time) {
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

	tkn := models.RefreshToken{
		Token: tokenString,
	}

	query := repositories.InsertInto(tkn.Name(), tkn, ctx)
	result, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("CreationRefreshTokenFailed")))
	}
	user.ID, _ = result.LastInsertId()

	return tokenString, expirationTime
}

func (obj *tokenService) GetRefreshToken(db *sql.DB, ctx context.Context, token string) *models.RefreshToken {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	refreshToken := &models.RefreshToken{}
	query := repositories.Select(refreshToken.Name(), map[string]any{
		"token": token,
	}, ctx)
	err := db.QueryRowContext(ctx, query).Scan(refreshToken)
	if err != nil {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("UsedRefreshToken")))
	}

	return refreshToken
}

func (obj *tokenService) SafeGetRefreshToken(db *sql.DB, ctx context.Context, token string) *models.RefreshToken {
	refreshToken := &models.RefreshToken{}
	query := repositories.Select(refreshToken.Name(), map[string]any{
		"token": token,
	}, ctx)
	err := db.QueryRowContext(ctx, query).Scan(refreshToken)
	if err != nil {
		return nil
	}

	return refreshToken
}

func (obj *tokenService) DeleteRefreshToken(db *sql.DB, ctx context.Context, token string) {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	query := repositories.Delete(models.RefreshTokenName, map[string]any{
		"token": token,
	}, ctx)
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(errors.New(errors.UnexpectedStatus, errors.ReSignIn, translate("DeletionRefreshTokenFailed")))
	}
}

func New() contract.TokenService {
	return instance
}
