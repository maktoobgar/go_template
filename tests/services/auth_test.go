package tests

import (
	"testing"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	service_auth "github.com/maktoobgar/go_template/internal/services/auth"
	"github.com/maktoobgar/go_template/tests"
)

func TestAuth(t *testing.T) {
	db := tests.New()
	auth := service_auth.New()

	username, password := "maktoobgar1", "123456789"

	// this should pass
	_, err := auth.SignUp(db, username, password, username)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignIn(db, username, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// duplicate error has to happen here
	_, err = auth.SignUp(db, username, password, username)
	if err == nil {
		t.Errorf("here had to happend an error because data is duplicate")
		return
	}

	// signing again, this should pass
	_, err = auth.SignIn(db, username, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignUp(db, "b", "444", "b")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignUp(db, "k", "444", "k")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// signing in another user, this should pass
	_, err = auth.SignIn(db, "k", "444")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
