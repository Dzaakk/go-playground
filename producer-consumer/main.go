package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(id int, jobs chan<- int, wg *sync.WaitGroup, producerWg *sync.WaitGroup) {
	defer wg.Done()
	defer producerWg.Done()

	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+100))

		job := i*100 + id

		fmt.Printf("Producer %d created job %d\n", id, job)
		jobs <- job
	}

	fmt.Printf("Producer %d finished\n", id)
}

func consumer(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Consumer %d started job %d\n", id, job)

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+500))

		fmt.Printf("Consumer %d finished job %d\n", id, job)
	}

	fmt.Printf("Consumer %d finished \n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan int, 10)

	var wg, producerWg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		producerWg.Add(1)
		go producer(i, jobs, &wg, &producerWg)
	}

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(i, jobs, &wg)
	}

	go func() {
		producerWg.Wait()
		fmt.Println("All producers finished, closing the jobs channel")
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("All work completed")
}
