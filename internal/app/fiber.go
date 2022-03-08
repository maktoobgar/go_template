package app

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/routers"
	"github.com/maktoobgar/go_template/pkg/errors"
)

var ()

func Fiber() {
	// Print Info
	info()

	// Config Server
	app := fiber.New(fiber.Config{
		Prefork:      true,
		ErrorHandler: errors.ErrorHandler,
		ReadTimeout:  time.Duration(time.Second * 5),
		WriteTimeout: time.Duration(time.Second * 30),
		IdleTimeout:  time.Duration(time.Minute * 5),
		AppName:      "Brand New App",
	})

	// Make App Global
	g.App = app

	// Router Settings
	routers.SetDefaultSettings(app)
	routers.Http(app)
	routers.Ws(app)

	// Run App
	g.Log().Panic(app.Listen(fmt.Sprintf("%s:%s", g.CFG.Api.IP, g.CFG.Api.Port)).Error(), Fiber, nil)
}
