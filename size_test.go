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

func TestSize_SizeOf_int(t *testing.T) {
	var v int
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int")
	}
}

func TestSize_SizeOf_int8(t *testing.T) {
	var v int8
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int8")
	}
}

func TestSize_SizeOf_int16(t *testing.T) {
	var v int16
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int16")
	}
}

func TestSize_SizeOf_int32(t *testing.T) {
	var v int32
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int32")
	}
}

func TestSize_SizeOf_int64(t *testing.T) {
	var v int64
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for int64")
	}
}

func TestSize_SizeOf_uint(t *testing.T) {
	var v uint
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uint")
	}
}

func TestSize_SizeOf_uint8(t *testing.T) {
	var v uint8
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uint8")
	}
}

func TestSize_SizeOf_uint16(t *testing.T) {
	var v uint16
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uint16")
	}
}

func TestSize_SizeOf_uint32(t *testing.T) {
	var v uint32
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uint32")
	}
}

func TestSize_SizeOf_uint64(t *testing.T) {
	var v uint64
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uint64")
	}
}

func TestSize_SizeOf_uintptr(t *testing.T) {
	var v uintptr
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for uintptr")
	}
}

func TestSize_SizeOf_float32(t *testing.T) {
	var v float32
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for float32")
	}
}

func TestSize_SizeOf_float64(t *testing.T) {
	var v float64
	size, err := SizeOf(v)
	if err != nil {
		t.Error(err)
	}
	if size != unsafe.Sizeof(v) {
		t.Error("invalid size for float64")
	}
}

//var c64 complex64
//var c128 complex128
//var iarr [42]int
//var ichan chan int
//var itf interface{}
//var imap map[int]int
//var iptr *int
//var islice []int
//var str string
//var srt struct{}
