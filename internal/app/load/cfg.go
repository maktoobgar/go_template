package app

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	iconfig "github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/internal/databases"
	g "github.com/maktoobgar/go_template/internal/global"
	session_service "github.com/maktoobgar/go_template/internal/services/session"
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
	setPwd()
	if err := config.ReadProjectConfigs(cfg.PWD, cfg); err != nil {
		log.Fatalln(err)
	}

	if err := config.ReadLocalConfigs(cfg.PWD, cfg); err != nil {
		log.Fatalln(err)
	}

	config.SetConfig(cfg)
	g.CFG = cfg
}

func initializeSession() {
	g.Session = session.New(session.Config{
		Expiration:   (time.Hour * 24) * 7,
		Storage:      session_service.New(),
		KeyLookup:    "header:session_id",
		CookieSecure: !g.CFG.Debug,
		CookieDomain: g.CFG.Domain,
	})
}

// Translator initialization
func initialTranslator() {
	t, err := translator.New(filepath.Join(cfg.PWD, cfg.Translator.Path), languages[0], languages[1:]...)
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

// Run dbs
func intialDBs() {
	err := databases.Setup(cfg)
	if err != nil {
		log.Fatalln(err)
	}
}

// Server initialization
func init() {
	initializeConfigs()
	initializeSession()
	initialTranslator()
	initialLogger()
	intialDBs()
}
