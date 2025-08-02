package main

// stage 1: generate numbers and send to a channel
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()

	return out
}

// stage 2: receives numbers, squares them, send to the next channel
func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()

	return out
}

// stage 3: receives numbers, adds 5, send to the final channel
func add(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n + 5:
			case <-done:
				return
			}
		}
	}()

	return out
}

// func main() {
// 	done := make(chan struct{})
// 	defer close(done)

// 	//setup the pipelin
// 	numbers := []int{1, 2, 3, 4, 5}
// 	c1 := generator(done, numbers...)
// 	c2 := square(done, c1)
// 	c3 := add(done, c2)

// 	//consume final output
// 	for res := range c3 {
// 		fmt.Println(res)
// 	}
// }
