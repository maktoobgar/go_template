package routers

import (
	"net/http"

	httpHandler "github.com/maktoobgar/go_template/internal/handlers/http"
	m "github.com/maktoobgar/go_template/internal/middleware"
)

func basicMiddlewares(next http.Handler, methods ...string) http.Handler {
	return m.Panic(m.Json(m.Translator(m.Cors(m.Method(next, methods...)))))
}

func HTTP(mux *http.ServeMux) {
	mux.Handle("/", basicMiddlewares(httpHandler.NotFound))
	mux.Handle("/api/", basicMiddlewares(httpHandler.Hi, "GET"))
	mux.Handle("/api/me", basicMiddlewares(m.Auth(httpHandler.Me), "GET"))

	// /api/auth
	{
		mux.Handle("/api/auth/sign_in", basicMiddlewares(httpHandler.SignIn, "POST"))
		mux.Handle("/api/auth/sign_up", basicMiddlewares(httpHandler.SignUp, "POST"))
		mux.Handle("/api/auth/refresh", basicMiddlewares(httpHandler.Refresh, "POST"))
	}
}
