package mmm

import (
	"testing"
	"unsafe"
)

// -----------------------------------------------------------------------------

func TestBytes_BytesOf_bool(t *testing.T) {
	var v bool = true
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*bool)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for bool")
	}
}

func TestBytes_BytesOf_int(t *testing.T) {
	var v int = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*int)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for int")
	}
}

func TestBytes_BytesOf_int8(t *testing.T) {
	var v int8 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*int8)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for int8")
	}
}

func TestBytes_BytesOf_int16(t *testing.T) {
	var v int16 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*int16)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for int16")
	}
}

func TestBytes_BytesOf_int32(t *testing.T) {
	var v int32 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*int32)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for int32")
	}
}

func TestBytes_BytesOf_int64(t *testing.T) {
	var v int64 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*int64)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for int64")
	}
}

func TestBytes_BytesOf_uint(t *testing.T) {
	var v uint = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uint)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uint")
	}
}

func TestBytes_BytesOf_uint8(t *testing.T) {
	var v uint8 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uint8)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uint8")
	}
}

func TestBytes_BytesOf_uint16(t *testing.T) {
	var v uint16 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uint16)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uint16")
	}
}

func TestBytes_BytesOf_uint32(t *testing.T) {
	var v uint32 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uint32)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uint32")
	}
}

func TestBytes_BytesOf_uint64(t *testing.T) {
	var v uint64 = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uint64)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uint64")
	}
}

func TestBytes_BytesOf_uintptr(t *testing.T) {
	var v uintptr = 42
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*uintptr)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for uintptr")
	}
}

func TestBytes_BytesOf_float32(t *testing.T) {
	var v float32 = 42.7
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*float32)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for float32")
	}
}

func TestBytes_BytesOf_float64(t *testing.T) {
	var v float64 = 42.7
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*float64)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for float64")
	}
}

func TestBytes_BytesOf_complex64(t *testing.T) {
	var v complex64 = 42.7
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*complex64)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for complex64")
	}
}

func TestBytes_BytesOf_complex128(t *testing.T) {
	var v complex128 = 42.7
	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*complex128)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for complex128")
	}
}

func TestBytes_BytesOf_array_int_64(t *testing.T) {
	var v [64]int
	for i := range v {
		v[i] = i
	}

	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*[64]int)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for [64]int")
	}
}

func TestBytes_BytesOf_array_int_312(t *testing.T) {
	var v [312]int
	for i := range v {
		v[i] = i
	}

	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*[312]int)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for [312]int")
	}
}

func TestBytes_BytesOf_array_int_array_27_518(t *testing.T) {
	var v [27][518]int
	for i := range v {
		for j := range v[i] {
			v[i][j] = i + i*j
		}
	}

	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*[27][518]int)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for [27][518]int")
	}
}

func TestBytes_BytesOf_struct_small(t *testing.T) {
	type small struct {
		a, b complex128
		c    int8
		z    [13]bool
	}
	v := small{a: 66.6, b: 42.7, c: 3, z: [13]bool{true}}

	size := unsafe.Sizeof(v)
	bytes := make([]byte, size)

	if err := BytesOf(v, bytes); err != nil {
		t.Error(err)
	}

	rv := *((*small)(unsafe.Pointer(&(bytes[0]))))
	if v != rv {
		t.Error("invalid bytes for struct{small}")
	}
}
