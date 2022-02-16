package app

import (
	"log"

	iconfig "github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/internal/databases"
	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/config"
	"github.com/maktoobgar/go_template/pkg/logging"
	"github.com/maktoobgar/go_template/pkg/translator"
	"golang.org/x/text/language"
)

var (
	cfg       = &iconfig.Config{}
	languages = []language.Tag{language.English, language.Persian}
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

// Translator initialization
func initialTranslator() {
	t, err := translator.New(cfg.Translator.Path, languages[0], languages[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	g.Translator = t.(*translator.TranslatorPack)
}

// Logger initialization
func initialLogger() {
	k := cfg.Logging
	opt := logging.Option(k)
	l, err := logging.New(&opt)
	if err != nil {
		log.Fatalln(err)
	}
	g.Logger = l.(*logging.LogBundle)
}

func intialDBs() {
	// Run dbs
	err := databases.Setup(cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

// Server initialization
func init() {
	initializeConfigs()
	initialTranslator()
	initialLogger()
	intialDBs()
}
