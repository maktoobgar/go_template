package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

var allowed_methods = []string{
	"GET", "POST", "PUT", "PATCH", "DELETE",
}

func Method(next http.Handler, methods ...string) http.Handler {
	if len(methods) == 0 {
		methods = allowed_methods
	}
	for i := range methods {
		methods[i] = strings.ToUpper(methods[i])
	}

	for i := range methods {
		found := false
		for j := range allowed_methods {
			if methods[i] == allowed_methods[j] {
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("%s is not allowed", methods[i])
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		translate := ctx.Value("translate").(translator.TranslatorFunc)
		currentMethod := strings.ToUpper(r.Method)
		found := false
		for i := range methods {
			if methods[i] == currentMethod {
				found = true
				break
			}
		}
		if !found {
			panic(errors.New(errors.MethodNotAllowedStatus, errors.DoNothing, translate("MethodNotAllowed")))
		}
		next.ServeHTTP(w, r)
	})
}
