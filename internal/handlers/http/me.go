package httpHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/models"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	c.JSON(user.Clean())
	return nil
}
