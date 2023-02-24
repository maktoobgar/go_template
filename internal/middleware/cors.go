package middleware

import (
	"fmt"
	"net/http"
	"strings"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

var allow_headers = "Origin, Content-Length, Content-Type"
var allow_methods = "GET, POST, PUT, PATCH, DELETE, OPTIONS"

func AddHeaders(headers []string) {
	for i := 0; i < len(headers); i++ {
		allow_headers += ", " + headers[i]
	}
}

func Cors(next http.Handler) http.Handler {
	allow_origins := strings.Split(g.CFG.AllowOrigins, ",")

	for i := range allow_origins {
		allow_origins[i] = strings.TrimSpace(allow_origins[i])
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		translate := ctx.Value("translate").(translator.TranslatorFunc)

		// Check if origin exists
		// Otherwise it is not a cors request
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}

		// request is not a CORS request but have origin header.
		// for example, use fetch api
		host := r.Host
		if origin == "http://"+host || origin == "https://"+host {
			next.ServeHTTP(w, r)
			return
		}

		// Check for origin access
		found := false
		for i := range allow_origins {
			if allow_origins[i] == origin {
				found = true
				break
			}
		}

		// Forbid if origin does not match
		if !found && len(allow_origins) != 0 && allow_origins[0] != "*" {
			panic(errors.New(errors.ForbiddenStatus, errors.DoNothing, translate("CorsError")))
		} else if strings.ToUpper(r.Method) == "OPTIONS" {
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Max-Age", fmt.Sprint(g.CFG.MaxAge))
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", allow_methods)
			w.Header().Set("Access-Control-Allow-Headers", allow_headers)
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
