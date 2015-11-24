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
