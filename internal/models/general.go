package models

import (
	"reflect"
	"strings"
)

func clean(obj interface{}, not interface{}, except ...string) map[string]interface{} {
	t := reflect.TypeOf(obj).Elem()
	tv := reflect.ValueOf(obj).Elem()
	nt := reflect.TypeOf(not).Elem()
	ntv := reflect.ValueOf(not).Elem()

	out := map[string]interface{}{}
	for i := 0; i < tv.NumField(); i++ {
		add := true
		for j := 0; j < ntv.NumField(); j++ {
			if !add {
				break
			}
			if t.Field(i).Name == nt.Field(j).Name {
				add = false
			}
		}
		for k := 0; k < len(except); k++ {
			if add {
				break
			}
			if strings.EqualFold(t.Field(i).Name, except[k]) {
				add = true
			}
		}
		if add {
			out[t.Field(i).Name] = tv.Field(i).Interface()
		}
	}

	return out
}
