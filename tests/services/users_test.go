package tests

import (
	"fmt"
	"testing"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	user_service "github.com/maktoobgar/go_template/internal/services/users"
	"github.com/maktoobgar/go_template/tests"
)

func TestUser(t *testing.T) {
	db := tests.New()
	uService := user_service.New()

	// user should create
	user, err := uService.CreateUser(db, "maktoobgar", "123456789", "maktoobgar")
	if err != nil {
		t.Errorf("error on creating a user, err: %s", err.Error())
		return
	}

	// user should not create
	_, err = uService.CreateUser(db, "maktoobgar", "123456789", "maktoobgar")
	if err == nil {
		t.Errorf("no error received on creating a duplicate user")
		return
	}

	// user should return
	_, err = uService.GetUser(db, user.Username)
	if err != nil {
		t.Errorf("error on getting user based on username, err: %s", err.Error())
		return
	}

	// user should return
	_, err = uService.GetUserByID(db, fmt.Sprint(user.ID))
	if err != nil {
		t.Errorf("error on getting user based on id, err: %s", err.Error())
		return
	}

	// passwords should match
	if !uService.CheckPasswordHash("123456789", user.Password) {
		t.Errorf("provided passwords should match but didn't")
		return
	}

	// passwords should not match
	if uService.CheckPasswordHash("1", user.Password) {
		t.Errorf("provided passwords should not match but did")
		return
	}

	// passwords should match
	if !uService.CheckPasswordHash("1", uService.HashPassword("1")) {
		t.Errorf("provided passwords should match but didn't")
		return
	}

	// passwords should not match
	if uService.CheckPasswordHash("12", uService.HashPassword("1")) {
		t.Errorf("provided passwords should not match but did")
		return
	}
}
