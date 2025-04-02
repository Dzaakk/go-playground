package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// race main
// func main() {
// 	// This is our shared variable
// 	counter := 0

// 	// We'll use this to wait for all goroutines to finish
// 	var wg sync.WaitGroup

// 	// Launch 1000 goroutines
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			// Each goroutine increments the counter
// 			counter++
// 			wg.Done()
// 		}()
// 	}

// 	// Wait for all goroutines to complete
// 	wg.Wait()

// 	// Print the final value
// 	fmt.Println("Final counter value:", counter)
// }

// fixed with mutex
// func main() {
// 	// This is our shared variable
// 	counter := 0

// 	// We'll use this to wait for all goroutines to finish
// 	var wg sync.WaitGroup

// 	// This mutex will protect our counter
// 	var mu sync.Mutex

// 	// Launch 1000 goroutines
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			// Lock the mutex before accessing the counter
// 			mu.Lock()
// 			counter++
// 			// Unlock when done
// 			mu.Unlock()

// 			wg.Done()
// 		}()
// 	}

// 	// Wait for all goroutines to complete
// 	wg.Wait()

// 	// Print the final value
// 	fmt.Println("Final counter value:", counter)
// }

// fixed with atomic operations
func main() {
	var counter int64 = 0

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Final counter value:", counter)
}
