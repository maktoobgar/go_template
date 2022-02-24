package tests

import (
	"testing"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	token_service "github.com/maktoobgar/go_template/internal/services/token"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/tests"
)

func TestToken(t *testing.T) {
	db := tests.New()
	tService := token_service.New()
	uService := user_service.New()

	// user should create
	user, err := uService.CreateUser(db, "maktoobgar2", "123456789", "maktoobgar")
	if err != nil {
		t.Errorf("error on creating a user for testing tokens, err:\n%s", err.Error())
		return
	}

	// token should create
	_, _, err = tService.CreateAccessToken(user)
	if err != nil {
		t.Errorf("error on creating access token, err: %s", err.Error())
		return
	}

	// token should create
	refreshToken, _, err := tService.CreateRefreshToken(db, user)
	if err != nil {
		t.Errorf("error on creating refresh token, err: %s", err.Error())
		return
	}

	// token should delete
	err = tService.DeleteRefreshToken(db, refreshToken)
	if err != nil {
		t.Errorf("error on deleting refresh token, err: %s", err.Error())
		return
	}

	// token should not delete
	err = tService.DeleteRefreshToken(db, refreshToken)
	if err != nil {
		t.Errorf("error on deleting pre deleted refresh token, err: %s", err.Error())
		return
	}

	// token should create
	refreshToken, _, err = tService.CreateRefreshToken(db, user)
	if err != nil {
		t.Errorf("error on creating refresh token, err: %s", err.Error())
		return
	}

	// token should receive
	_, err = tService.GetRefreshToken(db, refreshToken)
	if err != nil {
		t.Errorf("error on receiving refresh token from database, err: %s", err.Error())
		return
	}

	// token should delete
	err = tService.DeleteRefreshToken(db, refreshToken)
	if err != nil {
		t.Errorf("error on deleting refresh token, err: %s", err.Error())
		return
	}

	// token should not receive
	_, err = tService.GetRefreshToken(db, refreshToken)
	if err == nil {
		t.Errorf("no error on receiving a pre deleted refresh token out of database")
		return
	}
}
