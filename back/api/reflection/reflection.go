package reflection

import (
	"reflect"
	"runtime"
)

func GetFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
