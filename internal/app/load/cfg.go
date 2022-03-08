package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2/middleware/session"
	iconfig "github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/internal/databases"
	g "github.com/maktoobgar/go_template/internal/global"
	csrf_service "github.com/maktoobgar/go_template/internal/services/csrf"
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
	g.SecretKey = []byte(cfg.SecretKey)
}

// Initialization for session_service
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
func initialDBs() {
	err := databases.Setup(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	var db *goqu.Database = nil
	var ok bool = false
	if !g.CFG.Debug {
		db, ok = g.AllDBs["auth"]
		if !ok {
			db, ok = g.AllDBs["main"]
			if !ok {
				log.Fatalln(errors.New("both 'main' and 'auth' dbs are not defined (at least one of them required)"))
			}
		}
	} else {
		db, ok = g.AllDBs["test"]
		if !ok {
			log.Fatalln(errors.New("'test' db is not defined"))
		}
	}

	csrf_service.SetDB(db)
	session_service.SetDB(db)
}

// Server initialization
func init() {
	initializeConfigs()
	initializeSession()
	initialTranslator()
	initialLogger()
	initialDBs()
}
