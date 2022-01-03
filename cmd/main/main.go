package main

import (
	"log"

	"github.com/gotorn/core/internal/app"
	"github.com/gotorn/core/internal/config"
)

var (
	cfg = &config.Config{}
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
