package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/maktoobgar/go_template/pkg/errors"
)

type PanicResponse struct {
	Message string `json:"message"`
	Action  int    `json:"action"`
	Code    int    `json:"code"`
}

func Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			errInterface := recover()
			if errInterface == nil {
				return
			}
			err := errInterface.(error)
			code, action, message := errors.HttpError(err)
			res := PanicResponse{
				Message: message,
				Action:  action,
				Code:    code,
			}
			resBytes, _ := json.Marshal(res)
			w.WriteHeader(code)
			w.Write(resBytes)
		}()
		next.ServeHTTP(w, r)
	})
}
