package mmm

import (
	"fmt"
	"reflect"
	"unsafe"
)

// -----------------------------------------------------------------------------

// BytesOf writes the in-memory representation of `v` in `bytes`, including the
// content of its pointers.
func BytesOf(v interface{}, bytes []byte) error {
	return bytesOf(v, bytes)
}

func bytesOf(v interface{}, bytes []byte) error {
	t := reflect.TypeOf(v)
	k := t.Kind()
	switch k {
	case reflect.Bool:
		rv := v.(bool)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Int:
		rv := v.(int)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Int8:
		rv := v.(int8)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Int16:
		rv := v.(int16)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Int32:
		rv := v.(int32)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Int64:
		rv := v.(int64)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uint:
		rv := v.(uint)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uint8:
		rv := v.(uint8)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uint16:
		rv := v.(uint16)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uint32:
		rv := v.(uint32)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uint64:
		rv := v.(uint64)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Uintptr:
		rv := v.(uintptr)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Float32:
		rv := v.(float32)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Float64:
		rv := v.(float64)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Complex64:
		rv := v.(complex64)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Complex128:
		rv := v.(complex128)
		b := *((*[unsafe.Sizeof(rv)]byte)(unsafe.Pointer(&rv)))
		copy(bytes, b[:])
	case reflect.Array:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Chan:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Func:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Interface:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Map:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Ptr:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Slice:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.String:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.Struct:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	case reflect.UnsafePointer:
		return Error(fmt.Sprintf("type not supported (yet?): %s", k))
	default:
		return Error(fmt.Sprintf("`v` is not sizable (?): %#v", v))
	}

	return nil
}
