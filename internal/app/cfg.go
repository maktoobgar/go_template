package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/maktoobgar/go_template/build"
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

// Set Project PWD
func setPwd() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for parent := pwd; true; parent = filepath.Dir(parent) {
		if _, err := os.Stat(filepath.Join(parent, "go.mod")); err == nil {
			cfg.PWD = parent
			break
		}
	}
	os.Chdir(cfg.PWD)
}

// Initialization for config files in configs folder
func initializeConfigs() {
	// Loads default config, you just have to hard code it
	if err := config.ParseYamlBytes(build.Config, cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadLocalConfigs(cfg); err != nil {
		log.Fatalln(err)
	}

	g.CFG = cfg
	g.SecretKey = []byte(cfg.SecretKey)
}

// Translator initialization
func initialTranslator() {
	t, err := translator.New(build.Translations, languages[0], languages[1:]...)
	if err != nil {
		log.Fatalln(err)
	}
	g.Translator = t
}

// Logger initialization
func initialLogger() {
	k := cfg.Logging
	opt := logging.Option(k)
	l, err := logging.New(&opt)
	if err != nil {
		log.Fatalln(err)
	}
	g.Logger = l
}

// Run dbs
func initialDBs() {
	err := databases.Setup(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	var ok bool = false
	if !g.CFG.Debug {
		_, ok = g.AllDBs["main"]
		if !ok {
			log.Fatalln(errors.New("'main' db is not defined (required)"))
		}
	} else {
		_, ok = g.AllDBs["test"]
		if !ok {
			log.Fatalln(errors.New("'test' db is not defined"))
		}
	}
}

// Server initialization
func init() {
	setPwd()
	initializeConfigs()
	initialTranslator()
	initialLogger()
	initialDBs()
}
