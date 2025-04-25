package main

func containsDuplicate(nums []int) bool {
	check := make(map[int]bool, len(nums))
	if len(nums) < 2 {
		return false
	}
	for _, num := range nums {
		if check[num] {
			return true
		}
		check[num] = true
	}

	return false
}
