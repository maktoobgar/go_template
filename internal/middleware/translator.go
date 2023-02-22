package middleware

import (
	"context"
	"net/http"

	g "github.com/maktoobgar/go_template/internal/global"
)

func Translator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		lang := r.Header.Get("Accept-Language")
		if lang == "" {
			lang = "en"
		}
		ctx = context.WithValue(ctx, "translate", g.Translator.TranslateFunction(lang))
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
