package httpHandlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
)

func Me(c *fiber.Ctx) error {
	uService := user_service.New()

	session, err := g.Session.Get(c)
	if err != nil {
		return err
	}

	if session.Fresh() {
		session.Destroy()
		c.Redirect("SignIn")
		return nil
	}

	userID := session.Get(session.ID()).(int)
	user, err := uService.GetUserByID(fmt.Sprint(userID))
	if err != nil {
		return err
	}

	c.JSON(user.Clean())
	return nil
}
