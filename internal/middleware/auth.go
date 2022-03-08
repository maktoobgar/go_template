package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
)

func sessionAuth(c *fiber.Ctx) error {
	uService := user_service.New()
	session, err := g.Session.Get(c)
	if err != nil {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("Unauthorized"))
	}

	id := session.Get(session.ID()).(int)
	user, err := uService.GetUserByID(g.DB, fmt.Sprint(id))
	if err != nil {
		return err
	}

	c.Locals("user", user)
	return c.Next()
}

func tokenAuth(c *fiber.Ctx, token string) error {
	uService := user_service.New()
	claims := &contract.Claims{}

	// Token validation checks
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("InvalidToken"))
		}
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("InvalidToken"))
	}
	if !tkn.Valid {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("InvalidToken"))
	}

	// Check token is not refresh token
	if claims.Type != contract.AccessTokenType {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("NotAccessToken"))
	}

	user, err := uService.GetUser(g.DB, claims.Username)
	if err != nil {
		return err
	}

	c.Locals("user", user)
	return c.Next()
}

func Auth(c *fiber.Ctx) error {
	// Checking for session authentication
	sessionID, _ := g.Session.Storage.Get(c.GetReqHeaders()["Session_id"])
	if cookieSessionID := c.Request().Header.Cookie("session_id"); sessionID == nil {
		c.GetReqHeaders()["Session_id"] = string(cookieSessionID)
	}
	if sessionID != nil {
		return sessionAuth(c)
	}

	// Checking for token authentication
	var token = c.GetReqHeaders()["Token"]
	if cookieToken := string(c.Request().Header.Cookie("token")); token == "" {
		token = cookieToken
	}
	if token != "" {
		return tokenAuth(c, token)
	}

	return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Trans().TranslateEN("Unauthorized"))
}
