package main

// O(n) O(n)
// func majorityElement(nums []int) int {
// 	mapNum := map[int]int{}
// 	var count, e int
// 	for _, n := range nums {
// 		mapNum[n]++

// 		if count < mapNum[n] {
// 			count = mapNum[n]
// 			e = n
// 		}
// 	}
// 	return e
// }

func majorityElement(nums []int) int {
	count := 0
	candidate := 0

	for _, n := range nums {
		if count == 0 {
			candidate = n
		}

		if n == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}
