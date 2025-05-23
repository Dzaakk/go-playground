package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)

}

func sumArray(arr []int, index int) int {
	if index >= len(arr) {
		return 0
	}

	return arr[index] + sumArray(arr, index+1)
}

func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] > target {
		return binarySearch(arr, target, left, mid-1)
	}

	return binarySearch(arr, target, mid+1, right)
}

func printFileSystem(path string, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += " "
	}
	fmt.Println(indent + path)

	if level < 2 {
		children := []string{}

		if level == 0 {
			children = []string{"documents", "pictures", "downloads"}
		} else if path == "documents" {
			children = []string{"work", "personal"}
		} else if path == "pictures" {
			children = []string{"vacation", "family"}
		}

		for _, child := range children {
			printFileSystem(child, level+1)
		}
	}
}

func main() {
	// Example 1: Factorial calculation
	n := 5
	fmt.Printf("Factorial of %d is %d\n", n, factorial(n))

	// Example 2: Sum of array elements
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Sum of array elements: %d\n", sumArray(arr, 0))

	// Example 3: Binary search
	sortedArray := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
	target := 23
	index := binarySearch(sortedArray, target, 0, len(sortedArray)-1)
	fmt.Printf("Element %d found at index: %d\n", target, index)

	// Example 4: Directory traversal simulation
	printFileSystem("root", 0)

	// // Example 5: Tower of Hanoi
	// fmt.Println("\nTower of Hanoi solution for 3 disks:")
	// towerOfHanoi(3, "A", "C", "B")

	// // Example 6: Greatest Common Divisor (GCD)
	// a, b := 48, 18
	// fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd(a, b))

	// // Example 7: String palindrome check
	// testStr := "racecar"
	// fmt.Printf("Is '%s' a palindrome? %t\n", testStr, isPalindrome(testStr, 0, len(testStr)-1))
}
