package httpHandlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/maktoobgar/go_template/internal/contract"
	g "github.com/maktoobgar/go_template/internal/global"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/pkg/errors"
)

func Refresh(c *fiber.Ctx) error {
	tService := token_service.New()

	// Checking for token authentication
	var token = c.GetReqHeaders()["Token"]
	if cookieToken := string(c.Request().Header.Cookie("token")); token == "" {
		token = cookieToken
	}
	if token == "" {
		return errors.New(errors.InvalidStatus, errors.ReSingIn, g.Translator.TranslateEN("NotIncludedToken"))
	}

	uService := user_service.New()
	claims := &contract.Claims{}

	// Token validation
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return g.SecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("InvalidToken"))
		}
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("Unauthorized"))
	}
	if !tkn.Valid {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("Unauthorized"))
	}
	if claims.Type != contract.RefreshTokenType {
		return errors.New(errors.UnauthorizedStatus, errors.ReSingIn, g.Translator.TranslateEN("NotRefreshToken"))
	}

	// Check if refresh token is not used before
	_, err = tService.GetRefreshToken(token)
	if err != nil {
		return err
	}

	// Get user object
	user, err := uService.GetUser(claims.Username)
	if err != nil {
		return err
	}

	// Generate access and refresh tokens
	tokenString, expirationTime, err := tService.CreateAccessToken(user)
	if err != nil {
		return err
	}
	refreshTokenString, _, err := tService.CreateRefreshToken(user)
	if err != nil {
		return err
	}

	err = tService.DeleteRefreshToken(token)
	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	data := map[string]string{}
	data["AccessToken"] = tokenString
	data["RefreshToken"] = refreshTokenString
	return c.JSON(data)
}
