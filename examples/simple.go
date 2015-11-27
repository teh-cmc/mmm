// Copyright © 2015 Clement 'cmc' Rey <cr.rey.clement@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

/////
// Simple example of usage
//
//   go run examples/simple.go
//
/////

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/teh-cmc/mmm"
)

type Coordinate struct {
	x, y int
}

func main() {
	// create a new memory chunk that contains 3 Coordinate structures
	mc, err := mmm.NewMemChunk(Coordinate{}, 3)
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
}
