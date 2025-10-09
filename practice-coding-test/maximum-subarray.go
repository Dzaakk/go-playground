package main

func maxSubArray(nums []int) int {
	var tmp, maxSum int

	for _, n := range nums {
		if tmp+n > n {
			tmp += n
		} else {
			tmp = n
		}

		if tmp > maxSum {
			maxSum = tmp
		}
	}

	return maxSum
}
