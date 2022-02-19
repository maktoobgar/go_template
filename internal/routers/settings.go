package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	g "github.com/maktoobgar/go_template/internal/global"
	csrf_service "github.com/maktoobgar/go_template/internal/services/csrf"
)

func SetDefaultSettings(app *fiber.App) {
	app.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
		}),
		csrf.New(csrf.Config{
			Next:         csrf_service.Next,
			KeyLookup:    "form:csrf",
			CookieName:   "csrf",
			CookieSecure: !g.CFG.Debug,
			CookieDomain: g.CFG.Domain,
			Storage:      csrf_service.New(),
			ContextKey:   "csrf",
		}),
	)
}
