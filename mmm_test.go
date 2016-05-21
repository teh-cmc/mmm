// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"unsafe"
)

// -----------------------------------------------------------------------------

type Coordinate struct {
	x, y int
}

func Example_simple_usage() {
	// create a new memory chunk that contains 3 Coordinate structures
	mc, err := NewMemChunk(Coordinate{}, 3)
	if err != nil {
		log.Fatal(err)
	}

	// print 3
	fmt.Println(mc.NbObjects())

	// write {3,9} at index 0, then print {3,9}
	fmt.Println(mc.Write(0, Coordinate{3, 9}))
	// write {17,2} at index 1, then print {17,2}
	fmt.Println(mc.Write(1, Coordinate{17, 2}))
	// write {42,42} at index 2, then print {42,42}
	fmt.Println(mc.Write(2, Coordinate{42, 42}))

	// print {17,2}
	fmt.Println(mc.Read(1))
	// print {42,42}
	fmt.Println(*((*Coordinate)(unsafe.Pointer(mc.Pointer(2)))))

	// free memory chunk
	if err := mc.Delete(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// 3
	// {3 9}
	// {17 2}
	// {42 42}
	// {17 2}
	// {42 42}
}

// TestNewMemChunk tests the NewMemChunk function.
func TestNewMemChunk(t *testing.T) {
	// can't create 0-length chunk
	_, err := NewMemChunk(nil, 0)
	if err.Error() != "`n` must be > 0" {
		t.Error("expected length error, got", err)
	}

	types := []interface{}{
		// interface
		interface{}(0),
		// boolean
		false,
		// numeric
		byte(0),
		int(0),
		uint(0),
		uintptr(0),
		// array
		[3]int{},
		// unsafe pointer
		unsafe.Pointer(new(int)),
	}
	for _, typ := range types {
		// create chunk of type
		mc, err := NewMemChunk(typ, 1)
		if err != nil {
			t.Fatal(err)
		}
		// Read should produce input value
		if mc.Read(0) != typ {
			t.Error(typ)
		}
		mc.Delete()
	}

	invalidTypes := []interface{}{
		// nil pointer
		nil,
		// non-nil pointer
		new(int),
	}
	for _, typ := range invalidTypes {
		// create chunk of type
		_, err := NewMemChunk(typ, 1)
		if err == nil {
			t.Fatal("expected error, got nil")
		}
	}

	// run GC cycle; finalizers should run
	runtime.GC()
}
