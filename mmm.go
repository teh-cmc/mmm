// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"encoding/binary"
	"reflect"
	"syscall"
	"unsafe"
)

// -----------------------------------------------------------------------------

// Error represents an error within the mmm package.
type Error string

// Error implements the built-in error interface.
func (e Error) Error() string {
	return string(e)
}

// -----------------------------------------------------------------------------

// MemChunk represents a chunk of manually allocated memory.
type MemChunk struct {
	chunkSize uintptr
	objSize   uintptr

	itf       interface{}
	byteOrder binary.ByteOrder
	bytes     []byte
}

// NbObjects returns the number of objects in the chunk.
func (mc MemChunk) NbObjects() uint {
	return uint(mc.chunkSize / mc.objSize)
}

// -----------------------------------------------------------------------------

// Read returns the i-th object of the chunk as an interface.
//
// mmm doesn't provide synchronization of reads and writes on a MemChunk: it's
// entirely up to you to decide how you want to manage thread-safety.
//
// This will panic if `i` is out of bounds.
func (mc MemChunk) Read(i int) interface{} {
	// build a zero-value interface similar to the original one
	itf := reflect.Zero(reflect.ValueOf(mc.itf).Type()).Interface()
	// get a pointer to the memory representation of the new interface
	itfBytes := ((*[unsafe.Sizeof(itf)]byte)(unsafe.Pointer(&itf)))
	// ignore the bytes of the interface that correspond to the data
	ptrSize := unsafe.Sizeof(uintptr(0))
	itfLen := uintptr(len(itfBytes)) - ptrSize

	// get a pointer to the data bytes corresponding to index `i`
	dataPtr := uintptr(unsafe.Pointer(&(mc.bytes[uintptr(i)*mc.objSize])))

	// replace the data bytes of the interface with the data at index `i`
	switch ptrSize {
	case unsafe.Sizeof(uint8(0)):
		itfBytes[itfLen] = byte(dataPtr)
	case unsafe.Sizeof(uint16(0)):
		mc.byteOrder.PutUint16(itfBytes[itfLen:], uint16(dataPtr))
	case unsafe.Sizeof(uint32(0)):
		mc.byteOrder.PutUint32(itfBytes[itfLen:], uint32(dataPtr))
	case unsafe.Sizeof(uint64(0)):
		mc.byteOrder.PutUint64(itfBytes[itfLen:], uint64(dataPtr))
	}

	return itf
}

// Write writes the passed value to the i-th object of the chunk.
//
// It returns the passed value.
//
// mmm doesn't provide synchronization of reads and writes on a MemChunk: it's
// entirely up to you to decide how you want to manage thread-safety.
//
// This will panic if `i` is out of bounds, or if `v` is of a different type than
// the other objects in the chunk. Or if anything went wrong.
func (mc *MemChunk) Write(i int, v interface{}) interface{} {
	// panic if `v` is of a different type
	if reflect.TypeOf(v) != reflect.TypeOf(mc.itf) {
		panic("illegal value")
	}
	// copies `v`'s byte representation to index `i`
	if err := BytesOf(v, mc.bytes[uintptr(i)*mc.objSize:]); err != nil {
		panic(err)
	}

	return v
}

// Pointer returns a pointer to the i-th object of the chunk.
//
// It returns uintptr instead of unsafe.Pointer so that code using mmm
// cannot obtain unsafe.Pointers without importing the unsafe package
// explicitly.
//
// This will panic if `i` is out of bounds.
func (mc MemChunk) Pointer(i int) uintptr {
	return uintptr(unsafe.Pointer(&(mc.bytes[uintptr(i)*mc.objSize])))
}

// -----------------------------------------------------------------------------

// NewMemChunk returns a new memory chunk.
//
// Supported types:
// interfaces,
// arrays,
// structs,
// numerics and boolean (bool/int/uint/float/complex and their variants),
// unsafe.Pointer,
// and any possible combination of the above.
//
// `v`'s memory representation will be used as a template for the newly
// allocated memory. All data will be copied.
// `n` is the number of `v`-like objects the memory chunk can contain (i.e.,
// sizeof(chunk) = sizeof(v) * n).
func NewMemChunk(v interface{}, n uint) (MemChunk, error) {
	if n == 0 {
		return MemChunk{}, Error("`n` must be > 0")
	}

	size := reflect.ValueOf(v).Type().Size()
	bytes, err := syscall.Mmap(
		0, 0, int(size*uintptr(n)),
		syscall.PROT_READ|syscall.PROT_WRITE,
		syscall.MAP_PRIVATE|syscall.MAP_ANONYMOUS,
	)
	if err != nil {
		return MemChunk{}, err
	}

	for i := uint(0); i < n; i++ {
		if err := BytesOf(v, bytes[i*uint(size):]); err != nil {
			return MemChunk{}, err
		}
	}

	return MemChunk{
		chunkSize: size * uintptr(n),
		objSize:   size,

		itf:       v,
		byteOrder: Endianness(),
		bytes:     bytes,
	}, nil
}

// Delete frees the memory chunk.
func (mc *MemChunk) Delete() error {
	err := syscall.Munmap(mc.bytes)
	if err != nil {
		return err
	}

	mc.chunkSize = 0
	mc.objSize = 1
	mc.bytes = nil

	return nil
}

// -----------------------------------------------------------------------------

// Endianness returns the byte order of the current architecture.
func Endianness() binary.ByteOrder {
	var byteOrder binary.ByteOrder = binary.LittleEndian
	var i int = 0x1
	if ((*[unsafe.Sizeof(uintptr(0))]byte)(unsafe.Pointer(&i)))[0] == 0 {
		byteOrder = binary.BigEndian
	}

	return byteOrder
}
