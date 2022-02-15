package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/middleware"
	"github.com/maktoobgar/go_template/internal/services"
)

func AddRoutes(app *fiber.App) {
	app.Group("/api", middleware.Useless).Get(":name?", services.Hi).Name("Hi")
}
