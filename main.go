package main

import (
	"flag"

	"github.com/maktoobgar/go_template/internal/app"
)

var (
	whichApp = flag.String("app", "both", "defines which app to run")
)

func main() {
	app.API()
}
