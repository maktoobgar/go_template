package app

import (
	"fmt"

	"github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/internal/routers"
	"github.com/maktoobgar/go_template/pkg/api"
	"github.com/maktoobgar/go_template/pkg/logging"
	"github.com/maktoobgar/go_template/pkg/translator"
	"golang.org/x/text/language"
)

var (
	languages = []language.Tag{language.English, language.Persian}
)

func Run(cfg *config.Config) error {
	// Translator initialization
	translator, err := translator.New(cfg.Translator.Path, languages[0], languages[1:]...)
	if err != nil {
		return err
	}
	_ = translator

	// Logger initialization
	k := cfg.Logging
	l := logging.Option(k)
	logger, err := logging.New(&l)
	if err != nil {
		return err
	}
	_ = logger

	// Run App
	app := api.New("Brand New Game")
	routers.SetDefaultSettings(app)
	routers.AddRoutes(app)
	logger.Panic(app.Listen(fmt.Sprintf("%s:%s", cfg.Api.IP, cfg.Api.Port)).Error(), Run, nil)

	return nil
}
