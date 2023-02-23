package httpHandlers

import (
	"encoding/json"
	"net/http"

	"github.com/golodash/galidator"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/handlers/utils"
	"github.com/maktoobgar/go_template/internal/models"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/translator"
)

type signUpRequest struct {
	Username    string `json:"username" g:"required"`
	Password    string `json:"password" g:"required"`
	DisplayName string `json:"display_name" g:"required"`
}

type signUpResponse struct {
	User         models.UserCore `json:"user"`
	AccessToken  string          `json:"access_token"`
	RefreshToken string          `json:"refresh_token"`
}

var (
	signUpValidator = generator.Validator(signUpRequest{}, galidator.Messages{"required": "$field is required"})
)

func signUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	translate := ctx.Value("translate").(translator.TranslatorFunc)
	uService := user_service.New()
	tService := token_service.New()
	req := &signUpRequest{}
	utils.ParseBody(r.Body, translate, req)
	utils.ValidateBody(req, signUpValidator, translate)

	user := uService.CreateUser(g.DB, ctx, req.Username, req.Password, req.DisplayName)

	tokenString, _ := tService.CreateAccessToken(user, ctx)

	refreshTokenString, _ := tService.CreateRefreshToken(g.DB, ctx, user)

	res := signUpResponse{
		User:         user.UserCore,
		AccessToken:  tokenString,
		RefreshToken: refreshTokenString,
	}
	resBytes, _ := json.Marshal(res)
	w.Write(resBytes)
}

var SignUp = g.Handler{
	Handler: signUp,
}
