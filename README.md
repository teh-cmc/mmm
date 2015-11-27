# mmm ![Status](https://img.shields.io/badge/status-stable-green.svg?style=plastic) [![Build Status](http://img.shields.io/travis/teh-cmc/mmm.svg?style=plastic)](https://travis-ci.org/teh-cmc/mmm) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=plastic)](http://godoc.org/github.com/teh-cmc/mmm)

Manual memory management for golang.

## What you should know

Go doesn't provide any manual memory management primitives. [**For very good reasons**]().
This has been talked about numerous times on the [go-nuts mailing list](), have a look over there for detailed discussions.

**Unless you are absolutely certain that you have no better alternative, please do not use this library.**

`mmm` is no black magic: it simply allocates memory segments outside of the GC-managed heap and provides a clean API to abstract away all of the evil stuff that's really going on behind the scenes.

The performances of Go's garbage collector depend heavily on the number of pointers in your software.
*No matter how much performance you gain from using `mmm`, you could have had the same performance gains had you redesigned your software to avoid the use of pointers entirely.*

This is the raison d'etre of `mmm`: in some cases, purposefully (re)designing your software to avoid the use of pointers leads to code that is overly complex, harder to reason about, and thus, harder to maintain. In such cases, `mmm` might allow you to completely eliminate any GC overhead you're facing while still keeping your original design (with minimal changes to your implementation, of course).

`mmm` heavily relies on Go's implementation of interfaces.

## Install

```bash
go get -u github.com/teh-cmc/mmm
```

## Example

Here's a simple example of usage (code [here](examples/simple.go)):

```Go
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
```

## Demonstration

## License ![License](https://img.shields.io/badge/license-MIT-blue.svg?style=plastic)

The MIT License (MIT) - see LICENSE for more details

Copyright (c) 2015  Clement 'cmc' Rey  <cr.rey.clement@gmail.com>
