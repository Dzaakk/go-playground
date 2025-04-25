package main

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)

	for i, num := range nums {
		difference := target - num

		if targetIdx, found := hashMap[difference]; found {
			return []int{targetIdx, i}
		} else {
			hashMap[num] = i
		}
	}

	return []int{-1, -1}
}
