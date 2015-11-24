package mmm

import (
	"fmt"
	"reflect"
	"unsafe"
)

// -----------------------------------------------------------------------------

// BytesOf writes the in-memory representation of `v` in `bytes`, including the
// content of its pointers (i.e. they are flattened).
func BytesOf(v interface{}, bytes []byte) error {
	return bytesOf(v, bytes)
}

func bytesOf(v interface{}, bytes []byte) error {
	t := reflect.TypeOf(v)
	k := t.Kind()
	switch k {
	case reflect.Bool:
		return bytesOfNumericType(v, bytes)
	case reflect.Int:
		return bytesOfNumericType(v, bytes)
	case reflect.Int8:
		return bytesOfNumericType(v, bytes)
	case reflect.Int16:
		return bytesOfNumericType(v, bytes)
	case reflect.Int32:
		return bytesOfNumericType(v, bytes)
	case reflect.Int64:
		return bytesOfNumericType(v, bytes)
	case reflect.Uint:
		return bytesOfNumericType(v, bytes)
	case reflect.Uint8:
		return bytesOfNumericType(v, bytes)
	case reflect.Uint16:
		return bytesOfNumericType(v, bytes)
	case reflect.Uint32:
		return bytesOfNumericType(v, bytes)
	case reflect.Uint64:
		return bytesOfNumericType(v, bytes)
	case reflect.Uintptr:
		return bytesOfNumericType(v, bytes)
	case reflect.Float32:
		return bytesOfNumericType(v, bytes)
	case reflect.Float64:
		return bytesOfNumericType(v, bytes)
	case reflect.Complex64:
		return bytesOfNumericType(v, bytes)
	case reflect.Complex128:
		return bytesOfNumericType(v, bytes)
	case reflect.Array:
		return bytesOfArrayType(v, bytes)
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

func bytesOfNumericType(v interface{}, bytes []byte) error {
	vbytes := *((*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v)))
	vuintptr := *(*uintptr)(unsafe.Pointer(&(vbytes[unsafe.Sizeof(v)-unsafe.Sizeof(uintptr(0))])))
	// numeric type can't take more than 16 bytes (seriously it can't, right..?)
	vvalue := *((*[16]byte)(unsafe.Pointer(vuintptr)))
	copy(bytes, vvalue[:reflect.ValueOf(v).Type().Size()])

	return nil
}

func bytesOfArrayType(v interface{}, bytes []byte) error {
	vbytes := *((*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v)))
	vuintptr := *(*uintptr)(unsafe.Pointer(&(vbytes[unsafe.Sizeof(v)-unsafe.Sizeof(uintptr(0))])))

	size := reflect.ValueOf(v).Type().Size()
	for i := uintptr(0); size > 0; i += 128 {
		vvalue := *((*[128]byte)(unsafe.Pointer(vuintptr + i)))
		if size > 128 {
			copy(bytes[i:], vvalue[:128])
			size -= 128
		} else {
			copy(bytes[i:], vvalue[:size])
			size = 0
		}
	}

	return nil
}
