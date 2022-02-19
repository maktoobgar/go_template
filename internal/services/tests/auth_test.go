package tests

import (
	"os/exec"
	"testing"

	service_auth "github.com/maktoobgar/go_template/internal/services/auth"
)

func TestAuth(t *testing.T) {
	e1 := exec.Command("sql-migrate", "down", "-limit=0").Run()
	e2 := exec.Command("sql-migrate", "up").Run()
	if e1 != nil || e2 != nil {
		t.Error(e1, e2)
		return
	}

	auth := service_auth.New()
	username, password := "maktoobgar", "123456789"

	// this should pass
	_, err := auth.SignUp(username, password, username)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignIn(username, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// duplicate error has to happen here
	_, err = auth.SignUp(username, password, username)
	if err == nil {
		t.Errorf("here had to happend an error because data is duplicate")
		return
	}

	// signing again, this should pass
	_, err = auth.SignIn(username, password)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignUp("b", "444", "b")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// this should pass
	_, err = auth.SignUp("k", "444", "k")
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// signing in another user, this should pass
	_, err = auth.SignIn("k", "444")
	if err != nil {
		t.Errorf(err.Error())
		return
	}
}
