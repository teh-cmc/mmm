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
