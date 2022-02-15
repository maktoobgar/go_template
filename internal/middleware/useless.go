package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Useless(ctx *fiber.Ctx) error {
	fmt.Println("middleware")
	return ctx.Next()
}
