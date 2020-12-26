package gotprint

import (
	"fmt"
	"reflect"
)

func genericString(v reflect.Value) (result string) {
	switch v.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		s := fmt.Sprintf("%d", v.Int())
		return s

	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		s := fmt.Sprintf("%g", v.Float())
		return s

	case reflect.String:
		return v.String()
	}

	return ""
}
