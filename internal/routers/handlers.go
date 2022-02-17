package routers

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "github.com/maktoobgar/go_template/internal/handlers/http"
	"github.com/maktoobgar/go_template/internal/middleware"
)

func Http(app *fiber.App) {
	app.Group("/api", middleware.Useless).Get(":name?", httpHandler.Hi).Name("Hi")
}
