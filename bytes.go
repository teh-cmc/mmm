package mmm

import (
	"fmt"
	"reflect"
	"unsafe"
)

// -----------------------------------------------------------------------------

// BytesOf copies the in-memory representation of `v` in `bytes`.
//
// Only numeric, array, struct, pointer types and any combination of the above
// are supported.
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
	case reflect.Struct:
		return bytesOfStructType(v, bytes)
	case reflect.Ptr:
		return bytesOfPointerType(v, bytes)
	default:
		return Error(fmt.Sprintf("unsuppported type: %#v", k.String()))
	}

	return nil
}

// -----------------------------------------------------------------------------

func extractDataAddress(v interface{}) uintptr {
	// extract data address from the interface
	vbytes := *((*[unsafe.Sizeof(v)]byte)(unsafe.Pointer(&v)))
	vuintptr := *(*uintptr)(unsafe.Pointer(&(vbytes[unsafe.Sizeof(v)-unsafe.Sizeof(uintptr(0))])))

	return vuintptr
}

func bytesOfNumericType(v interface{}, bytes []byte) error {
	// interpret data as a contiguous array of 16 bytes
	vvalue := *((*[16]byte)(unsafe.Pointer(extractDataAddress(v))))
	// copy data byte-representation to `bytes`
	copy(bytes, vvalue[:reflect.ValueOf(v).Type().Size()])

	return nil
}

func bytesOfArrayType(v interface{}, bytes []byte) error {
	vuintptr := extractDataAddress(v)
	size := reflect.ValueOf(v).Type().Size()
	// loop until there are no more data chunks
	for i := uintptr(0); size > 0; i += 128 {
		// interpret data as a contiguous array of 128 bytes
		vvalue := *((*[128]byte)(unsafe.Pointer(vuintptr + i)))
		// copy data byte-representation to `bytes`
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

func bytesOfStructType(v interface{}, bytes []byte) error {
	return bytesOfArrayType(v, bytes)
}

func bytesOfPointerType(v interface{}, bytes []byte) error {
	return bytesOfNumericType(v, bytes)
}
