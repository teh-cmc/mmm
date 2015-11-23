package mmm

import (
	"testing"
	"unsafe"
)

// -----------------------------------------------------------------------------

func TestSize_SizeOf_bool(t *testing.T) {
	var v bool
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for bool")
	}
}
