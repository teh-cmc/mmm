// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"reflect"
	"unsafe"
)

// -----------------------------------------------------------------------------

// BytesOf copies the in-memory representation of `v` into `bytes`.
//
// Supported types:
// interfaces,
// arrays,
// structs,
// numerics and boolean (bool/int/uint/float/complex and their variants),
// unsafe.Pointer,
// and any possible combination of the above.
func BytesOf(v interface{}, bytes []byte) error {
	return bytesOf(v, bytes)
}

func bytesOf(v interface{}, bytes []byte) error {
	t, err := TypeOf(v)
	if err != nil {
		return err
	}

	switch t {
	case TypeNumeric, TypeArray, TypeStruct:
		// First extract v's data pointer and determine its size.
		vuintptr := (*[2]uintptr)(unsafe.Pointer(&v))[1]
		size := reflect.TypeOf(v).Size()

		// Construct a "fake slice" using the data pointer and size. This object's
		// memory layout will match that of []byte.
		slice := struct {
			data uintptr
			_len uintptr
			_cap uintptr
		}{vuintptr, size, size}

		// Convert the fake slice to a real byte slice and copy it.
		copy(bytes, *(*[]byte)(unsafe.Pointer(&slice)))

	case TypeUnsafePointer:
		// do nothing
	}

	return nil
}
