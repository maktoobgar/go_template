package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	wsHandlers "github.com/maktoobgar/go_template/internal/services/socket"
)

func Ws(app *fiber.App) {
	app.Get("/ws", websocket.New(wsHandlers.WS))
}
