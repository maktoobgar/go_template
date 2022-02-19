package httpHandlers

import (
	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	auth_service "github.com/maktoobgar/go_template/internal/services/auth"
)

type signInRequest struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
}

func SignIn(c *fiber.Ctx) error {
	req := &signInRequest{}
	if err := c.BodyParser(req); err != nil {
		return err
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
	c.JSON(data)
	return nil
}
