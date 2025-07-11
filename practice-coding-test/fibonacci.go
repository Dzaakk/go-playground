package main

func fibonacci(n int) []int {
	fibs := []int{}

	first := 0
	second := 1

	for i := 0; i < n; i++ {
		fibs = append(fibs, first)
		next := first + second
		first = second
		second = next
	}

	return fibs
}

// func main() {
// 	n := 10
// 	result := fibonacci(n)
// 	fmt.Println(result)
// }
