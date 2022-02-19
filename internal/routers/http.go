package routers

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "github.com/maktoobgar/go_template/internal/handlers/http"
	"github.com/maktoobgar/go_template/internal/middleware"
)

func Http(app *fiber.App) {
	// /api
	apiGroup := app.Group("/api", middleware.Useless)
	apiGroup.Get("me", httpHandler.Me).Name("Me")
	apiGroup.Get(":name?", httpHandler.Hi).Name("Hi")

	// /api/auth
	authGroup := apiGroup.Group("/auth")
	authGroup.Post("/signIn", httpHandler.SignIn).Name("SignIn")
}
