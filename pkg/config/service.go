package config

import (
	"reflect"
)

// Returns database username
func GetDBUsername() (res string) {
	c := reflect.ValueOf(cfg).Elem()

	defer func() {
		err := recover()
		if err != nil {
			res = ""
		}
	}()

	res = c.FieldByName("Database").FieldByName("Username").Interface().(string)
	return
}
