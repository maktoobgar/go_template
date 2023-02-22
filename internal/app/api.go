package app

import (
	"fmt"
	"net/http"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/routers"
)

func API() {
	// Print Info
	info()

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", g.CFG.Api.IP, g.CFG.Api.Port),
		Handler: mux,
	}

	g.Server = server

	// Router Settings
	routers.HTTP(mux)

	// Run App
	g.Logger.Panic(server.ListenAndServe().Error(), API, nil)
}
