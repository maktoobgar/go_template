package tests

import (
	"testing"
	"time"

	_ "github.com/maktoobgar/go_template/internal/app/load"
	csrf_service "github.com/maktoobgar/go_template/internal/services/csrf"
	"github.com/maktoobgar/go_template/tests"
)

func TestCSRF(t *testing.T) {
	csrf := csrf_service.New()
	csrf_service.SetDB(tests.New())

	value := []byte("asdawewrtergfgdmgkdrg4r5")
	key := "m"

	// data should set
	err := csrf.Set(key, value, time.Duration(time.Hour*24))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should have an duplication error
	err = csrf.Set(key, value, time.Duration(time.Hour*24))
	if err == nil {
		t.Errorf("we should have an error here because of the same data setting data happened")
		return
	}

	// we should get the data
	res, err := csrf.Get(key)
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
	err = csrf.Delete(key)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should get an error for requesting non existing data
	res, err = csrf.Get(key)
	if err == nil {
		t.Errorf("error has to heppen here but it is nil, err: %v", err)
		return
	}
	if res != nil {
		t.Errorf("in Get() after Delete() expected: %v, got: %v", nil, res)
		return
	}

	// add an expired key
	err = csrf.Set(key, value, time.Duration(-time.Hour*24))
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// we should get nothing here
	res, err = csrf.Get(key)
	if err == nil {
		t.Errorf("expire date was deprecated but data still returns, res: %v, err: %v", res, err)
		return
	}

	// we should delete everything here safely
	err = csrf.Reset()
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// just to say i tested everything
	csrf.Close()
}
