package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Useless(c *fiber.Ctx) error {
	fmt.Println("middleware")
	return c.Next()
}
