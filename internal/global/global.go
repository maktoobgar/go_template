package g

import (
	"database/sql"
	"net/http"

	"github.com/maktoobgar/go_template/internal/config"
	"github.com/maktoobgar/go_template/pkg/logging"
	"github.com/maktoobgar/go_template/pkg/translator"
)

// Handling section
type Handler struct {
	Handler func(w http.ResponseWriter, r *http.Request)
}

// Function that gets executed to host a url
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handler(w, r)
}

// Config
var CFG *config.Config = nil

// Utilities
var Logger logging.Logger = nil
var Translator translator.Translator = nil

// App
var Server *http.Server = nil

// AppSecret
var SecretKey []byte = nil

// Default DB
var DB *sql.DB = nil

// Connections
var PostgresCons = map[string]*sql.DB{}
var SqliteCons = map[string]*sql.DB{}
var MySQLCons = map[string]*sql.DB{}
var SqlServerCons = map[string]*sql.DB{}
var AllCons = map[string]*sql.DB{}
