package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

func worker(id int, jobs <-chan Job, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j.ID)
		//simulate work
		time.Sleep(time.Second)
		results <- fmt.Sprintf("Worker %d finished job %d with payload %s", id, j.ID, j.Payload)
	}
}

// func main() {
// 	const numJobs = 10
// 	const numWorkers = 3

// 	jobs := make(chan Job, numJobs)
// 	results := make(chan string, numJobs)

// 	var wg sync.WaitGroup

// 	for w := 1; w <= numWorkers; w++ {
// 		wg.Add(1)
// 		go worker(w, jobs, results, &wg)
// 	}

// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- Job{ID: j, Payload: "some-data"}
// 	}

// 	close(jobs)
// 	wg.Wait()

// 	for a := 1; a <= numJobs; a++ {
// 		fmt.Println(<-results)
// 	}
// }
