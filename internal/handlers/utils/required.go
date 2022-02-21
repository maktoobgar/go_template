package utils

import "reflect"

func Required(obj interface{}) bool {
	t := reflect.TypeOf(obj).Elem()
	v := reflect.ValueOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Tag.Get("required") == "true" && v.Field(i).Len() <= 0 {
			return false
		}
	}

	return true
}
