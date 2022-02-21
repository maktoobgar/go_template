package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
)

func SessionAuth(c *fiber.Ctx) error {
	fmt.Println("here")
	uService := user_service.New()
	_, err := g.Session.Storage.Get(c.GetReqHeaders()["session_id"])
	if err != nil {
		fmt.Println(err)
		c.Redirect("SignIn")
		return err
	}
	fmt.Println("here1")

	session, err := g.Session.Get(c)
	if err != nil {
		c.Redirect("SignIn")
		return err
	}
	fmt.Println("here2")

	id := session.Get(session.ID()).(int)
	user, err := uService.GetUserByID(fmt.Sprint(id))
	if err != nil {
		c.Redirect("SignIn")
		return err
	}
	fmt.Println("here3")

	c.Locals("user", user)
	return c.Next()
}
