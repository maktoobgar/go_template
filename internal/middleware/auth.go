package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func tokenAuth(token string, ctx context.Context) context.Context {
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	uService := user_service.New()
	claims := &contract.Claims{}

	// Token validation checks
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("InvalidToken")))
		}
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("InvalidToken")))
	}
	if !tkn.Valid {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("InvalidToken")))
	}

	// Check token is not refresh token
	if claims.Type != contract.AccessTokenType {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("NotAccessToken")))
	}

	user := uService.GetUser(g.DB, ctx, claims.Username)

	ctx = context.WithValue(ctx, "user", user)
	return ctx
}

// Checking for token authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		translate := ctx.Value("translate").(translator.TranslatorFunc)
		var token = r.Header.Get("Token")
		if token == "" {
			panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("Unauthorized")))
		}

		ctx = tokenAuth(token, r.Context())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
