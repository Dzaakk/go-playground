package main

import (
	"fmt"
	"time"
)

func SelectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Accept ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Accept ", msg2)
		}
	}
}
