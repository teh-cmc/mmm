# mmm ![Status](https://img.shields.io/badge/status-stable-green.svg?style=plastic) [![Build Status](http://img.shields.io/travis/teh-cmc/mmm.svg?style=plastic)](https://travis-ci.org/teh-cmc/mmm) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=plastic)](http://godoc.org/github.com/teh-cmc/mmm)

Manual memory management for golang.

Have a look at [FreeTree]() for a real-world example of how to use `mmm`.

## What you should know

Go doesn't provide any manual memory management primitives. [**For very good reasons**]().
This has been talked about numerous times on the [go-nuts mailing list](), have a look over there for detailed discussions.

To make it short: **unless you are absolutely certain that you have no better alternative and that you understand all of the tradoffs involved, please do not use this library.**

`mmm` is no black magic: it simply allocates memory segments outside of the GC-managed heap and provides a clean and simple API (`Read()`, `Write()`, `Pointer()`) that abstracts away all of the evil stuff that's acually going on behind the scenes.

The performances of Go's garbage collector depend heavily on the number of pointers your software is using.
*No matter how much performance you gain by using `mmm`, you could have had the same gains had you redesigned your software to avoid the use of pointers entirely.*

This is the raison d'etre of `mmm`: in some cases, purposefully (re)designing your software to avoid the use of pointers actually leads to code that is overly complex, harder to reason about, and thus, harder to maintain. In such cases, `mmm` might allow you to completely eliminate the GC overhead issues in your software while keeping your original design (with minimal changes to your implementation, of course).

Note that `mmm` heavily relies on Go's implementation of interfaces.

Finally, for the adventurous, you can find most the ugly stuff [here]() and [here]().

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
//	 go run examples/simple.go
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

Complete code for the following demonstration is available [here](experiment/experiment.go).

All of the results shown below were computed using a DELL XPS 15-9530 (i7-4712HQ@2.30GHz).

#### Case A: managed heap, 10 million pointers to int

Let's see what happens when we store 10 millions pointers to integer on the managed heap:

```Go
// build 10 million pointers to integer on the managed heap
ints := make([]*int, 10*1e6)
// init our pointers
for i := range ints {
	j := i
	ints[i] = &j
}

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected, and to prevent them from being optimized away
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(ints[i*1e4]))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (managed heap, 10 million pointers): %d us\n", (time.Now().UnixNano()-now)/1e3)
}
```

This prints:

```
value @ index 0: 0
GC time (managed heap, 10 million pointers): 329840 us
value @ index 10000: 10000
GC time (managed heap, 10 million pointers): 325375 us
value @ index 20000: 20000
GC time (managed heap, 10 million pointers): 323367 us
value @ index 30000: 30000
GC time (managed heap, 10 million pointers): 327905 us
value @ index 40000: 40000
GC time (managed heap, 10 million pointers): 326469 us
```

That's an average ~326ms per GC call.
Let's move to case B where we will start using `mmm`'s memory chunks.

#### Case B: unmanaged heap, pointers generated on-the-fly

`mmm` doesn't store any pointer, it doesn't need to.

Since the data is stored on an unmanaged heap, it cannot be collected even if there's no reference to it. This allows `mmm` to generate pointers only when something's actually reading or writing to the data.

In pratice, it looks like that:

```Go
// build 10 million integers on an unmanaged heap
intz, err := mmm.NewMemChunk(int(0), 10*1e6)
if err != nil {
	log.Fatal(err)
}
// init our integers
for i := 0; i < int(intz.NbObjects()); i++ {
	intz.Write(i, i)
}

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected (pointer to data is generated on-the-fly)
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, intz.Read(i*1e4))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (unmanaged heap, pointers generated on-the-fly): %d us\n", (time.Now().UnixNano()-now)/1e3)
```

This prints:

```
value @ index 0: 0
GC time (unmanaged heap, pointers generated on-the-fly): 999 us
value @ index 10000: 10000
GC time (unmanaged heap, pointers generated on-the-fly): 665 us
value @ index 20000: 20000
GC time (unmanaged heap, pointers generated on-the-fly): 827 us
value @ index 30000: 30000
GC time (unmanaged heap, pointers generated on-the-fly): 882 us
value @ index 40000: 40000
GC time (unmanaged heap, pointers generated on-the-fly): 1016 us
```

That's an average ~0.9ms per GC call.

We went from a ~326ms average to a ~0.9ms average; but the comparison isn't really fair now, is it? In case A we were storing every pointer, here we're simply not storing any.

That leads us to case C, in which we build pointers to each and every integer in our unamanaged heap.

#### Case C: unamanaged heap, storing all generated pointers

What happens when we build and store 10 million pointers for each and every integer in our unmanaged memory chunk?

```Go
// build 10 million unsafe pointers on the managed heap
ptrs := make([]unsafe.Pointer, 10*1e6)
// init those pointers so that they point to the unmanaged heap
for i := range ptrs {
	ptrs[i] = unsafe.Pointer(intz.Pointer(i))
}

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(*int)(ptrs[i*1e4]))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (unmanaged heap, all generated pointers stored): %d us\n", (time.Now().UnixNano()-now)/1e3)
}
```

This prints:

```
value @ index 0: 0
GC time (unmanaged heap, all generated pointers stored): 47196 us
value @ index 10000: 10000
GC time (unmanaged heap, all generated pointers stored): 47307 us
value @ index 20000: 20000
GC time (unmanaged heap, all generated pointers stored): 47485 us
value @ index 30000: 30000
GC time (unmanaged heap, all generated pointers stored): 47145 us
value @ index 40000: 40000
GC time (unmanaged heap, all generated pointers stored): 47221 us
```

The results here are pretty interesting: on the one hand this is ~47 times slower than case B (in which we used `mmm`'s memory chunks but didn't actually store any pointer) because unsafe pointers require far less work from the GC, but on the other hand this is still 6 times faster than case A (in which we used native pointers).

Six times faster is already quite the good deal, but why stop there? As we've already pointed out, `mmm` doesn't need to store references to its data... so, don't.

This is what case D is all about, in which we will convert those pointers into simple numeric references and store them as such.

#### Case D: unmanaged heap, storing numeric references

Instead of storing (unsafe) pointers, let's treat these pointers as what they really are: simple numeric references.

```Go
// build 10 million numeric references on the managed heap
refs := make([]uintptr, 10*1e6)
// init those references so that they each contain one of the addresses
for i := range refs {
	refs[i] = uintptr(ptrs[i])
}

// get rid of those unsafe pointers we stored
ptrs = nil

for i := 0; i < 5; i++ {
	// randomly print one of our integers to make sure it's all working
	// as expected
	fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(*int)(unsafe.Pointer(refs[i*1e4])))

	// run GC
	now := time.Now().UnixNano()
	runtime.GC()
	fmt.Printf("\tGC time (unmanaged heap, all numeric references stored): %d us\n", (time.Now().UnixNano()-now)/1e3)
}
```

This prints:

```
value @ index 0: 0
GC time (unmanaged heap, all numeric references stored): 715 us
value @ index 10000: 10000
GC time (unmanaged heap, all numeric references stored): 783 us
value @ index 20000: 20000
GC time (unmanaged heap, all numeric references stored): 882 us
value @ index 30000: 30000
GC time (unmanaged heap, all numeric references stored): 711 us
value @ index 40000: 40000
GC time (unmanaged heap, all numeric references stored): 723 us
```

We're basically back to the results of case B.
As far as the GC is concerned, those pointers don't exist, which translates into sub-millisecond GC calls.

Still, the memory they point to does exist, and is just one cast away from being read from and written to.

We now have everything we need to build pointer-based software without any GC overhead, and without any design modification: this is basically how [FreeTree]() is implemented.

## License ![License](https://img.shields.io/badge/license-MIT-blue.svg?style=plastic)

The MIT License (MIT) - see LICENSE for more details

Copyright (c) 2015	Clement 'cmc' Rey	<cr.rey.clement@gmail.com>
