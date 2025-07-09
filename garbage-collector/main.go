package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	fmt.Println("Before GC: ", m.Alloc)

	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Println("After GC: ", m.Alloc)
}
