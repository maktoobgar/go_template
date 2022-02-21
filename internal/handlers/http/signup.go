package httpHandlers

import (
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/handlers/utils"
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
		return errors.New(errors.InvalidStatus, "not all required fields provided")
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
