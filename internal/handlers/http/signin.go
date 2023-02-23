package httpHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/golodash/galidator"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/handlers/utils"
	"github.com/maktoobgar/go_template/internal/models"
	auth_service "github.com/maktoobgar/go_template/internal/services/auth"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	"github.com/maktoobgar/go_template/pkg/translator"
)

type signInRequest struct {
	Username string `json:"username" g:"required"`
	Password string `json:"password" g:"required"`
}

type signInResponse struct {
	User         models.UserCore `json:"user"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
}

var (
	signInValidator = generator.Validator(signInRequest{}, galidator.Messages{"required": "$field is required"})
)

func signIn(w http.ResponseWriter, r *http.Request) {
	tService := token_service.New()
	req := &signInRequest{}
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	utils.ParseBody(r.Body, translate, req)
	utils.ValidateBody(req, signInValidator, translate)

	auth := auth_service.New()
	user := auth.SignIn(g.DB, ctx, req.Username, req.Password)

	accessTokenString, _ := tService.CreateAccessToken(user, ctx)
	refreshTokenString, _ := tService.CreateRefreshToken(g.DB, ctx, user)

	res := signInResponse{
		User:         user.UserCore,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

var SignIn = g.Handler{
	Handler: signIn,
}
