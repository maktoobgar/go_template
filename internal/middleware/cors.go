package middleware

import (
	"net/http"
	"strings"

	g "github.com/maktoobgar/go_template/internal/global"
	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func Cors(next http.Handler) http.Handler {
	allow_origins := strings.Split(g.CFG.AllowOrigins, ",")
	for i := range allow_origins {
		allow_origins[i] = strings.TrimSpace(allow_origins[i])
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		translate := ctx.Value("translate").(translator.TranslatorFunc)
		origin := r.Header.Get("Origin")
		if origin == "" {
			next.ServeHTTP(w, r)
			return
		}
		found := false
		for i := range allow_origins {
			if allow_origins[i] == origin {
				found = true
				break
			}
		}
		if !found && len(allow_origins) != 0 && allow_origins[0] != "*" {
			panic(errors.New(errors.ForbiddenStatus, errors.DoNothing, translate("CorsError")))
		} else if strings.ToUpper(r.Method) == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "15")
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.WriteHeader(200)
			return
		}

		w.Header().Set("Access-Control-Max-Age", "15")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		next.ServeHTTP(w, r)
	})
}
