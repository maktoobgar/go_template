package routers

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "github.com/maktoobgar/go_template/internal/handlers/http"
	"github.com/maktoobgar/go_template/internal/middleware"
)

func Http(app *fiber.App) {
	// /api
	apiGroup := app.Group("/api", middleware.Useless)

	// /api/me
	meGroup := apiGroup.Group("/me", middleware.Auth)
	meGroup.Get("/", httpHandler.Me).Name("Me")

	// /api/:name?
	apiGroup.Get(":name?", httpHandler.Hi).Name("Hi")

	// /api/auth
	authGroup := apiGroup.Group("/auth")
	authGroup.Post("/signIn", httpHandler.SignIn).Name("SignIn")
	authGroup.Post("/signUp", httpHandler.SignUp).Name("SignUp")

	// /api/auth/token
	tokenGroup := authGroup.Group("/token")
	tokenGroup.Post("/signin", httpHandler.SignInToken).Name("SignInToken")
	tokenGroup.Post("/signup", httpHandler.SignUpToken).Name("SignUpToken")
	tokenGroup.Post("/refresh", httpHandler.Refresh).Name("Refresh")
}
