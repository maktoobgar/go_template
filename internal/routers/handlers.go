package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/middleware"
	httpHandler "github.com/maktoobgar/go_template/internal/services/http"
)

func Http(app *fiber.App) {
	app.Group("/api", middleware.Useless).Get(":name?", httpHandler.Hi).Name("Hi")
}
