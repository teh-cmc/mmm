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
