package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a := 0
	for b := 1; b < len(nums); b++ {
		if nums[a] != nums[b] {
			a++
			nums[a] = nums[b]
		}
	}

	return a + 1
}

// func main() {
// 	nums := []int{1, 1, 2}

// 	fmt.Println(removeDuplicates(nums))
// }
