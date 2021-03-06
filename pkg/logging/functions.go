package logging

import (
	"reflect"
	"runtime"
	"strings"
)

// Send a function as input param to this function and
// get the package name of that function as string
func getPackageName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	strs = strings.Split(strs[len(strs)-2], "/")
	return strs[len(strs)-1]
}

// Send a function as input param to this function and
// get the function name as string
func getFunctionName(temp interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

// Send a function as input param to this function and
// get the function Entry as string
//
// I have no use for this function in the code actually
// but I wrote it, and it works fine so let it be
func getFunctionEntry(temp interface{}) uintptr {
	return runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Entry()
}
