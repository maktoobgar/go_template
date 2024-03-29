package routes

import (
	"net/http"
	"time"

	g "github.com/maktoobgar/go_template/internal/global"
	httpHandler "github.com/maktoobgar/go_template/internal/handlers/http"
	m "github.com/maktoobgar/go_template/internal/middleware"
	"github.com/maktoobgar/go_template/pkg/router"
)

func basicMiddlewares(next http.Handler, methods ...string) http.Handler {
	return m.Translator(m.Panic(m.Timeout(time.Duration(g.CFG.Timeout), m.ConcurrentLimiter(g.CFG.MaxConcurrentRequests, m.Json(m.Cors(m.Method(next, methods...)))))))
}

func HTTP(mux *router.Router) {
	// /api
	{
		mux.Handle("/api/me/", basicMiddlewares(m.Auth(httpHandler.Me), "GET"))
		// /api/auth
		{
			mux.Handle("/api/auth/sign_in/", basicMiddlewares(httpHandler.SignIn, "POST"))
			mux.Handle("/api/auth/sign_up/", basicMiddlewares(httpHandler.SignUp, "POST"))
			mux.Handle("/api/auth/refresh/", basicMiddlewares(httpHandler.Refresh, "POST"))
		}
		mux.Handle("/api/.+/", basicMiddlewares(httpHandler.Hi, "GET"))
	}

	// Not found
	mux.Handle("/.*/", basicMiddlewares(httpHandler.NotFound, "GET"))
}
