package main

import (
	"flag"
	"strings"

	"github.com/maktoobgar/go_template/internal/app"
)

var (
	whichApp = flag.String("app", "fiber", "defines which app to run")
)

func main() {
	flag.Parse()
	w := strings.ToLower(*whichApp)
	if w == "fiber" || w == "f" || w == "0" {
		app.Fiber()
	} else if w == "grpc" || w == "g" || w == "1" {
		app.Grpc()
	}
}
