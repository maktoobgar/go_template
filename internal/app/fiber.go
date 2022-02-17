package app

import (
	"fmt"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/routers"
	"github.com/maktoobgar/go_template/pkg/api"
)

var ()

func Fiber() {
	// Run App
	app := api.New("Brand New App")
	g.App = app

	// Router Settings
	routers.SetDefaultSettings(app)
	routers.Http(app)
	routers.Ws(app)

	g.Log().Panic(app.Listen(fmt.Sprintf("%s:%s", g.CFG.Api.IP, g.CFG.Api.Port)).Error(), Fiber, nil)
}
