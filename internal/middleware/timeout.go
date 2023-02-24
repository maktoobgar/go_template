package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/maktoobgar/go_template/pkg/errors"
	"github.com/maktoobgar/go_template/pkg/translator"
)

func Timeout(timeout time.Duration, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		translate := r.Context().Value("translate").(translator.TranslatorFunc)
		// Create a new context with the given timeout
		ctx, cancel := context.WithTimeout(r.Context(), timeout*time.Second)
		defer cancel() // Cancel the context to release resources when the request completes

		done := make(chan struct{})
		go func() {
			Panic(next).ServeHTTP(w, r.WithContext(ctx))
			close(done)
		}()
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				panic(errors.New(errors.TimeoutStatus, errors.Resend, translate("TimeoutError")))
			}
		case <-done:
			return
		}
	})
}
