package main

import (
	"fmt"
	"time"
)

func workerSignalOnly(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	fmt.Println("Done!")
	done <- true
}

// func main() {
// 	done := make(chan bool)
// 	go worker(done)

// 	<-done
// }
