package tests

import (
	"testing"
	"time"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	session_service "github.com/maktoobgar/go_template/internal/services/session"
	"github.com/maktoobgar/go_template/tests"
)

func TestSession(t *testing.T) {
	session := session_service.New()
	session_service.SetDB(tests.New())

	value := []byte("asdawewrtergfgdmgkdrg4r5")
	key := "m"

	// data should set
	err := session.Set(key, value, time.Duration(time.Hour*24))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should have an duplication error
	err = session.Set(key, value, time.Duration(time.Hour*24))
	if err == nil {
		t.Errorf("we should have an error here because of the same data setting data happened")
		return
	}

	// we should get the data
	res, err := session.Get(key)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	for i := range res {
		if res[i] != value[i] {
			t.Errorf("in Get() expected: %v, got: %v", value, res)
			return
		}
	}

	// data should remove safely
	err = session.Delete(key)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should get an error for requesting non existing data
	res, err = session.Get(key)
	if err == nil {
		t.Errorf("error has to heppen here but it is nil, err: %v", err)
		return
	}
	if res != nil {
		t.Errorf("in Get() after Delete() expected: %v, got: %v", nil, res)
		return
	}

	// add an expired key
	err = session.Set(key, value, time.Duration(-time.Hour*24))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should get nothing here
	res, err = session.Get(key)
	if err == nil {
		t.Errorf("expire date was deprecated but data still return, res: %v, err: %v", res, err)
		return
	}

	// we should delete everything here safely
	err = session.Reset()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// just to say i tested everything
	session.Close()
}
