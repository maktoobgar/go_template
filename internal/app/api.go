package app

import (
	"fmt"
	"log"
	"net/http"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/internal/routes"
	"github.com/maktoobgar/go_template/pkg/router"
)

func API() {
	// Print Info
	info()

	mux := new(router.Router)
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", g.CFG.Api.IP, g.CFG.Api.Port),
		Handler: mux,
	}
	// Server uses ServeHTTP(ResponseWriter, *Request) method

	g.Server = server

	// Router Settings
	routes.HTTP(mux)

	// Run App
	log.Panic(server.ListenAndServe().Error())
}
