// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"testing"
	"unsafe"
)

// -----------------------------------------------------------------------------

func TestTypes_TypesOf(tt *testing.T) {
	var t Type
	var err error

	var b bool
	t, err = TypeOf(b)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var i int
	t, err = TypeOf(i)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var i8 int8
	t, err = TypeOf(i8)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var i16 int16
	t, err = TypeOf(i16)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var i32 int32
	t, err = TypeOf(i32)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var i64 int64
	t, err = TypeOf(i64)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ui uint
	t, err = TypeOf(ui)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ui8 uint8
	t, err = TypeOf(ui8)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ui16 uint16
	t, err = TypeOf(ui16)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ui32 uint32
	t, err = TypeOf(ui32)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ui64 uint64
	t, err = TypeOf(ui64)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var uiptr uintptr
	t, err = TypeOf(uiptr)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var f32 float32
	t, err = TypeOf(f32)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var f64 float64
	t, err = TypeOf(f64)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var c64 complex64
	t, err = TypeOf(c64)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var c128 complex128
	t, err = TypeOf(c128)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeNumeric {
		tt.Error("wrong type")
	}

	var ia [16]int
	t, err = TypeOf(ia)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeArray {
		tt.Error("wrong type")
	}

	var srt struct{}
	t, err = TypeOf(srt)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeStruct {
		tt.Error("wrong type")
	}

	var up unsafe.Pointer
	t, err = TypeOf(up)
	if err != nil {
		tt.Error(err)
	}
	if t != TypeUnsafePointer {
		tt.Error("wrong type")
	}

	var ic chan int
	t, err = TypeOf(ic)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var fn func()
	t, err = TypeOf(fn)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var itf interface{}
	t, err = TypeOf(itf)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var im map[int]int
	t, err = TypeOf(im)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var ip *int
	t, err = TypeOf(ip)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var is []int
	t, err = TypeOf(is)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}

	var str string
	t, err = TypeOf(str)
	if err == nil || t != TypeInvalid {
		tt.Error("illegal type")
	}
}
