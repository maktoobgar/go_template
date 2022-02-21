package httpHandlers

import (
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/handlers/utils"
	auth_service "github.com/maktoobgar/go_template/internal/services/auth"
	"github.com/maktoobgar/go_template/pkg/errors"
)

type signInRequest struct {
	Username string `json:"username" xml:"username" form:"username" required:"true"`
	Password string `json:"password" xml:"password" form:"password" required:"true"`
}

func SignIn(c *fiber.Ctx) error {
	req := &signInRequest{}
	if err := c.BodyParser(req); err != nil || !utils.Required(req) {
		return errors.New(errors.InvalidStatus, errors.Resend, g.Translator.TranslateEN("RequiresNotProvided"))
	}

	auth := auth_service.New()
	user, err := auth.SignIn(req.Username, req.Password)
	if err != nil {
		return err
	}

	session, err := g.Session.Get(c)
	if err != nil {
		return err
	}
	defer session.Save()

	if !session.Fresh() {
		err = session.Regenerate()
		if err != nil {
			return err
		}
	}
	session.Set(session.ID(), user.ID)

	data := user.Clean()
	data["SessionID"] = session.ID()
	return c.JSON(data)
}
