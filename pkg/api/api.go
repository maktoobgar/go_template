package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func New(appName string) *fiber.App {
	app := fiber.New(fiber.Config{
		// Prefork:      true,
		ReadTimeout:  time.Duration(time.Second * 5),
		WriteTimeout: time.Duration(time.Second * 30),
		IdleTimeout:  time.Duration(time.Minute * 5),
		AppName:      appName,
	})

	return app
}
