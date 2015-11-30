// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package mmm

import (
	"fmt"
	"log"
	"unsafe"
)

// -----------------------------------------------------------------------------

type Coordinate struct {
	x, y int
}

func Example() {
	// create a new memory chunk that contains 3 Coordinate structures
	mc, err := NewMemChunk(Coordinate{}, 3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mc.NbObjects())

	fmt.Println(mc.Write(0, Coordinate{3, 9}))
	fmt.Println(mc.Write(1, Coordinate{17, 2}))
	fmt.Println(mc.Write(2, Coordinate{42, 42}))

	fmt.Println(mc.Read(1))
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
