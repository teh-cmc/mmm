package main

import (
	"log"
	"math/rand"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/teh-cmc/mmm"
)

// -----------------------------------------------------------------------------

func main() {
	ints, err := mmm.NewMemChunk(42, 1e8)
	if err != nil {
		log.Fatal(err)
	}

	runtime.GC()
	debug.FreeOSMemory()

	for {
		log.Println(*(*int)(ints.Index(rand.Intn(int(ints.NbObjects())))))
		now := time.Now().UnixNano()
		runtime.GC()
		log.Printf("GC time: %d ms\n", (time.Now().UnixNano()-now)/1e6)
	}
}
