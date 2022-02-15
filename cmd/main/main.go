package main

import (
	"log"

	"github.com/maktoobgar/go_template/internal/app"
	iconfig "github.com/maktoobgar/go_template/internal/config"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/config"
)

var (
	cfg = &iconfig.Config{}
)

// Initialization for config files in configs folder
func initializeConfigs() {
	if err := config.ReadProjectConfigs(cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadLocalConfigs(cfg); err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(cfg)
	g.CFG = cfg
}

// Server initialization
func init() {
	initializeConfigs()
}

func main() {
	if err := app.Run(cfg); err != nil {
		log.Fatalln(err)
	}
}
