// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"reflect"
	"runtime"
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

	arr   reflect.Value
	bytes []byte
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
func (mc *MemChunk) Read(i int) interface{} {
	return mc.arr.Index(i).Interface()
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
	if reflect.TypeOf(v) != mc.arr.Type().Elem() {
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
	if _, err := TypeOf(v); err != nil {
		return MemChunk{}, err
	}

	t := reflect.TypeOf(v)
	size := t.Size()
	bytes, err := syscall.Mmap(
		0, 0, int(size*uintptr(n)),
		syscall.PROT_READ|syscall.PROT_WRITE,
		mmapFlags,
	)
	if err != nil {
		return MemChunk{}, err
	}

	for i := uintptr(0); i < uintptr(n); i++ {
		if err := BytesOf(v, bytes[i*size:]); err != nil {
			return MemChunk{}, err
		}
	}

	// create an array of type t, backed by the mmap'd memory
	arr := reflect.Zero(reflect.ArrayOf(int(n), t)).Interface()
	(*[2]uintptr)(unsafe.Pointer(&arr))[1] = uintptr(unsafe.Pointer(&bytes[0]))

	ret := MemChunk{
		chunkSize: size * uintptr(n),
		objSize:   size,

		arr:   reflect.ValueOf(arr),
		bytes: bytes,
	}

	// set a finalizer to free the chunk's memory when it would normally be
	// garbage collected
	runtime.SetFinalizer(&ret, func(chunk *MemChunk) {
		if chunk.bytes != nil {
			chunk.Delete()
		}
	})

	return ret, nil
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
