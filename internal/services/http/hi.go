package httpHandlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hi(ctx *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s", ctx.Params("name"))
	return ctx.SendString(msg)
}
