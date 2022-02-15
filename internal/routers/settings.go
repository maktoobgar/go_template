package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
)

func SetDefaultSettings(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}), csrf.New())
}
