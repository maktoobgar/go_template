package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
)

func SessionAuth(c *fiber.Ctx) error {
	uService := user_service.New()
	_, err := g.Session.Storage.Get(c.GetReqHeaders()["Session_id"])
	if err != nil {
		return err
	}

	session, err := g.Session.Get(c)
	if err != nil {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("Unauthorized"))
	}

	id := session.Get(session.ID()).(int)
	user, err := uService.GetUserByID(fmt.Sprint(id))
	if err != nil {
		return err
	}

	c.Locals("user", user)
	return c.Next()
}
