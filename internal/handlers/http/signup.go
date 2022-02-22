package httpHandlers

import (
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/handlers/utils"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type signUpRequest struct {
	Username    string `json:"username" xml:"username" form:"username" required:"true"`
	Password    string `json:"password" xml:"password" form:"password" required:"true"`
	DisplayName string `json:"display_name" xml:"display_name" form:"display_name" required:"true"`
}

func SignUp(c *fiber.Ctx) error {
	uService := user_service.New()
	req := &signUpRequest{}
	if err := c.BodyParser(req); err != nil || !utils.Required(req) {
		return errors.New(errors.InvalidStatus, errors.Resend, g.Translator.TranslateEN("RequiresNotProvided"))
	}

	user, err := uService.CreateUser(req.Username, req.Password, req.DisplayName)
	if err != nil {
		return err
	}

	session, err := g.Session.Get(c)
	if err != nil {
		return err
	}
	defer session.Save()

	session.Set(session.ID(), user.ID)

	data := user.Clean()
	data["SessionID"] = session.ID()
	return c.JSON(data)
}

func SignUpToken(c *fiber.Ctx) error {
	uService := user_service.New()
	tService := token_service.New()
	req := &signUpRequest{}
	if err := c.BodyParser(req); err != nil || !utils.Required(req) {
		return errors.New(errors.InvalidStatus, errors.Resend, g.Translator.TranslateEN("RequiresNotProvided"))
	}

	user, err := uService.CreateUser(req.Username, req.Password, req.DisplayName)
	if err != nil {
		return err
	}

	tokenString, expirationTime, err := tService.CreateAccessToken(user)
	if err != nil {
		return err
	}

	refreshTokenString, _, err := tService.CreateRefreshToken(user)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	data := user.Clean()
	data["AccessToken"] = tokenString
	data["RefreshToken"] = refreshTokenString
	return c.JSON(data)
}
