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

// Index returns the i-th object of the chunk as an interface.
//
// This will panic if `i` is out of bounds.
func (mc MemChunk) Index(i int) interface{} {
	itf := mc.itf
	itfBytes := ((*[unsafe.Sizeof(itf)]byte)(unsafe.Pointer(&itf)))

	dataPtr := uintptr(unsafe.Pointer(&(mc.bytes[uintptr(i)*mc.objSize])))

	ptrSize := unsafe.Sizeof(uintptr(0))
	itfLen := uintptr(len(itfBytes)) - ptrSize

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

// NewMemChunk returns a new memory chunk.
//
// `v`'s memory representation will be used as a template for the newly
// allocated memory. All pointers will be flattened. All data will be copied.
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

		itf:       reflect.ValueOf(v).Interface(),
		byteOrder: Endianness(),
		bytes:     bytes,
	}, nil
}

func Endianness() binary.ByteOrder {
	var byteOrder binary.ByteOrder = binary.LittleEndian
	var i int = 0x1
	if ((*[unsafe.Sizeof(uintptr(0))]byte)(unsafe.Pointer(&i)))[0] == 0 {
		byteOrder = binary.BigEndian
	}

	return byteOrder
}
