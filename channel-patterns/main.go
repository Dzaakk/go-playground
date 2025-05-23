package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
	}()

	val := <-ch
	fmt.Println(val)
}
