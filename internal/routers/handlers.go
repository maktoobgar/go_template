package routers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func AddRoutes(app *fiber.App) {
	// GET /api/register
	app.Get("/api/:name?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("name"))
		return c.SendString(msg)
	}).Name("temp")
}
