package app

import (
	"fmt"

	"github.com/maktoobgar/go_template/internal/config"
	g "github.com/maktoobgar/go_template/internal/global"
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
	t, err := translator.New(cfg.Translator.Path, languages[0], languages[1:]...)
	if err != nil {
		return err
	}
	g.Translator = t.(*translator.TranslatorPack)

	// Logger initialization
	k := cfg.Logging
	l := logging.Option(k)
	log, err := logging.New(&l)
	if err != nil {
		return err
	}
	g.Logger = log.(*logging.LogBundle)

	// Run Grpc
	go grpcRun()

	// Run App
	app := api.New("Brand New Game")
	g.App = app

	routers.SetDefaultSettings(app)
	routers.Http(app)
	routers.Ws(app)
	g.Log().Panic(app.Listen(fmt.Sprintf("%s:%s", cfg.Api.IP, cfg.Api.Port)).Error(), Run, nil)

	return nil
}
