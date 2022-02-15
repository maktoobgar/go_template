package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/maktoobgar/go_template/internal/services"
)

func WS(app *fiber.App) {
	app.Get("/ws", websocket.New(services.WS))
}
