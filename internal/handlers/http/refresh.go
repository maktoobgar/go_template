package httpHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

type refreshResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	tService := token_service.New()

	// Checking for token authentication
	var token = r.Header.Get("Token")
	if token == "" {
		panic(errors.New(errors.InvalidStatus, errors.ReSignIn, translate("NotIncludedToken")))
	}

	uService := user_service.New()
	claims := &contract.Claims{}

	// Token validation
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("InvalidToken")))
		}
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("Unauthorized")))
	}
	if !tkn.Valid {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("Unauthorized")))
	}
	if claims.Type != contract.RefreshTokenType {
		panic(errors.New(errors.UnauthorizedStatus, errors.ReSignIn, translate("NotRefreshToken")))
	}

	// Check if refresh token is not used before
	refreshToken := tService.SafeGetRefreshToken(g.DB, ctx, token)
	if refreshToken == nil {
		panic(errors.New(errors.InvalidStatus, errors.ReSignIn, translate("UsedRefreshToken")))
	}

	// Get user object
	user := uService.GetUser(g.DB, ctx, claims.Username)

	// Generate access and refresh tokens
	tokenString, _ := tService.CreateAccessToken(user, ctx)
	refreshTokenString, _ := tService.CreateRefreshToken(g.DB, ctx, user)

	tService.DeleteRefreshToken(g.DB, ctx, token)

	res := refreshResponse{
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

var Refresh = g.Handler{
	Handler: refresh,
}
