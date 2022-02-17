package session_service

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	g "github.com/maktoobgar/go_template/internal/global"
)

func init() {
	g.Session = session.New(session.Config{
		Expiration:   (time.Hour * 24) * 7,
		Storage:      New(),
		KeyLookup:    "header:session_id",
		CookieSecure: !g.CFG.Debug,
		CookieDomain: g.CFG.Domain,
	})
}
