package main

func workerPool(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// start 3 workers
	for i := 0; i < 3; i++ {
		go workerPool(jobs, results)
	}

	// send jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// collect results
	for i := 1; i <= 5; i++ {
		<-results
	}

}
