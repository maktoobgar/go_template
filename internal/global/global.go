package g

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/pkg/logging"
	"github.com/maktoobgar/go_template/pkg/translator"
)

// Config
var CFG *config.Config = nil

// Utilities
var Logger *logging.LogBundle = nil
var Translator *translator.TranslatorPack = nil

// App
var App *fiber.App = nil

// DBs
var Postgres = map[string]*goqu.Database{}
var Sqlite = map[string]*goqu.Database{}
var MySQL = map[string]*goqu.Database{}
var SqlServer = map[string]*goqu.Database{}

// Default DB
var DB *goqu.Database = nil

// Connections
var PostgresCons = map[string]*sql.DB{}
var SqliteCons = map[string]*sql.DB{}
var MySQLCons = map[string]*sql.DB{}
var SqlServerCons = map[string]*sql.DB{}

func Log() logging.Logger {
	return Logger
}

func Trans() translator.Translator {
	return Translator
}
