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

// BytesOf copies the in-memory representation of `v` in `bytes`.
//
// Only numeric, array, struct types and any combination of the above are
// supported.
func BytesOf(v interface{}, bytes []byte) error {
	return bytesOf(v, bytes)
}

func bytesOf(v interface{}, bytes []byte) error {
	t, err := TypeOf(v)
	if err != nil {
		return err
	}

	switch t {
	case TypeNumeric:
		return bytesOfNumericType(v, bytes)
	case TypeArray:
		return bytesOfArrayType(v, bytes)
	case TypeStruct:
		return bytesOfStructType(v, bytes)
	case TypeUnsafePointer:
		// do nothing
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
	if err := TypeCheck(v); err != nil {
		return err
	}

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
