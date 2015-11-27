package main

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"
	"time"
	"unsafe"

	"github.com/teh-cmc/mmm"
)

// -----------------------------------------------------------------------------

/////
// This file shows various ways of using mmm's MemChunk's and how they affect
// GC's performances compared to native Go pointers.
/////

func main() {

	/////////////////////////////////////////////////
	// A: Managed heap, 10 million pointers to int //
	/////////////////////////////////////////////////
	fmt.Println(`Case A: what happens when we store 10 million pointers to integer
on the managed heap?` + "\n")

	// build 10 million pointers to integer on the managed heap
	ints := make([]*int, 10*1e6)
	// init our pointers
	for i := range ints {
		j := i
		ints[i] = &j
	}

	// get rid of init garbage
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected, and to prevent them from being optimized away
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(ints[i*1e4]))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (managed heap, 10 million pointers): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (managed heap, 10 million pointers): 329840 us
	//   value @ index 10000: 10000
	//   GC time (managed heap, 10 million pointers): 325375 us
	//   value @ index 20000: 20000
	//   GC time (managed heap, 10 million pointers): 323367 us
	//   value @ index 30000: 30000
	//   GC time (managed heap, 10 million pointers): 327905 us
	//   value @ index 40000: 40000
	//   GC time (managed heap, 10 million pointers): 326469 us

	fmt.Println()

	//////////////////////////////////////////////////////
	// B: Unmanaged heap, pointers generated on-the-fly //
	//////////////////////////////////////////////////////
	fmt.Println(`Case B: mmm doesn't store any pointer, it doesn't need to.
Since the data is stored on an unmanaged heap, it cannot be collected
even if there's no reference to it. This allows mmm to generate the
pointers only when something's asking for some data.` + "\n")

	// build 10 million integers on an unmanaged heap
	intz, err := mmm.NewMemChunk(int(0), 10*1e6)
	if err != nil {
		log.Fatal(err)
	}
	// init our integers
	for i := 0; i < int(intz.NbObjects()); i++ {
		intz.Write(i, i)
	}

	// get rid of (almost all) previous garbage
	ints = nil
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected (pointer to data is generated on-the-fly)
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, intz.Read(i*1e4))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (unmanaged heap, pointers generated on-the-fly): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (unmanaged heap, pointers generated on-the-fly): 999 us
	//   value @ index 10000: 10000
	//   GC time (unmanaged heap, pointers generated on-the-fly): 665 us
	//   value @ index 20000: 20000
	//   GC time (unmanaged heap, pointers generated on-the-fly): 827 us
	//   value @ index 30000: 30000
	//   GC time (unmanaged heap, pointers generated on-the-fly): 882 us
	//   value @ index 40000: 40000
	//   GC time (unmanaged heap, pointers generated on-the-fly): 1016 us

	fmt.Println()

	///////////////////////////////////////////////////////
	// C: Unmanaged heap, storing all generated pointers //
	///////////////////////////////////////////////////////
	fmt.Println("Case C: what happens when we store all the generated pointers?\n")

	// build 10 million unsafe pointers on the managed heap
	ptrs := make([]unsafe.Pointer, 10*1e6)
	// init those pointers so that they point to the unmanaged heap
	for i := range ptrs {
		ptrs[i] = unsafe.Pointer(intz.Pointer(i))
	}

	// get rid of (almost all) previous garbage
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(*int)(ptrs[i*1e4]))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (unmanaged heap, all generated pointers stored): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (unmanaged heap, all generated pointers stored): 47196 us
	//   value @ index 10000: 10000
	//   GC time (unmanaged heap, all generated pointers stored): 47307 us
	//   value @ index 20000: 20000
	//   GC time (unmanaged heap, all generated pointers stored): 47485 us
	//   value @ index 30000: 30000
	//   GC time (unmanaged heap, all generated pointers stored): 47145 us
	//   value @ index 40000: 40000
	//   GC time (unmanaged heap, all generated pointers stored): 47221 us

	fmt.Println()

	///////////////////////////////////////////////////
	// D: Unmanaged heap, storing numeric references //
	///////////////////////////////////////////////////
	fmt.Println(`Case D: as case C showed, storing all the generated pointers to the
unmanaged heap is still order of magnitudes faster than storing pointers
to the managed heap.
Still, why keep pointer values when our data is not garbage-collectable?
What happens if we store all the generated pointers as numeric references?` + "\n")

	// build 10 million numeric references on the managed heap
	refs := make([]uintptr, 10*1e6)
	// init those references so that they "point" to the unmanaged heap
	for i := range refs {
		refs[i] = uintptr(ptrs[i])
	}

	// get rid of (almost all) previous garbage
	ptrs = nil
	runtime.GC()
	debug.FreeOSMemory()

	for i := 0; i < 5; i++ {
		// randomly print one of our integers to make sure it's all working
		// as expected
		fmt.Printf("\tvalue @ index %d: %d\n", i*1e4, *(*int)(unsafe.Pointer(refs[i*1e4])))

		// run GC
		now := time.Now().UnixNano()
		runtime.GC()
		fmt.Printf("\tGC time (unmanaged heap, all numeric references stored): %d us\n", (time.Now().UnixNano()-now)/1e3)
	}

	// Results:
	//   value @ index 0: 0
	//   GC time (unmanaged heap, all numeric references stored): 715 us
	//   value @ index 10000: 10000
	//   GC time (unmanaged heap, all numeric references stored): 783 us
	//   value @ index 20000: 20000
	//   GC time (unmanaged heap, all numeric references stored): 882 us
	//   value @ index 30000: 30000
	//   GC time (unmanaged heap, all numeric references stored): 711 us
	//   value @ index 40000: 40000
	//   GC time (unmanaged heap, all numeric references stored): 723 us
}
