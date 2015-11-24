package mmm

import (
	"fmt"
	"reflect"
)

// -----------------------------------------------------------------------------

// SizeOf returns the real size of `v` in memory, including the content of its
// pointers (i.e. they are flattened).
func SizeOf(v interface{}) (uintptr, error) {
	return sizeOf(v)
}

func sizeOf(v interface{}) (uintptr, error) {
	var total uintptr

	t := reflect.TypeOf(v)
	k := t.Kind()
	switch k {
	case reflect.Bool:
		total += t.Size()
	case reflect.Int:
		total += t.Size()
	case reflect.Int8:
		total += t.Size()
	case reflect.Int16:
		total += t.Size()
	case reflect.Int32:
		total += t.Size()
	case reflect.Int64:
		total += t.Size()
	case reflect.Uint:
		total += t.Size()
	case reflect.Uint8:
		total += t.Size()
	case reflect.Uint16:
		total += t.Size()
	case reflect.Uint32:
		total += t.Size()
	case reflect.Uint64:
		total += t.Size()
	case reflect.Uintptr:
		total += t.Size()
	case reflect.Float32:
		total += t.Size()
	case reflect.Float64:
		total += t.Size()
	case reflect.Complex64:
		total += t.Size()
	case reflect.Complex128:
		total += t.Size()
	case reflect.Array:
		total += t.Size()
	case reflect.Chan:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Func:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Interface:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Map:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Ptr:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Slice:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.String:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Struct:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.UnsafePointer:
		return 0, Error(fmt.Sprintf("type not supported (yet?): %s", k))
	default:
		return 0, Error(fmt.Sprintf("`v` is not sizable (?): %#v", v))
	}

	return total, nil
}
