package main

import (
	"fmt"
	"sync"
)

func race() {
	counter := 0

	for i := 0; i < 5; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
			}
		}()
	}

	fmt.Scanln()
	fmt.Println("Final Counter:", counter)
}

func notRace() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // Output should be 5000
}
