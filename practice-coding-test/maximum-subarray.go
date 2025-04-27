package main

func maxSubArray(nums []int) int {
	currSum := 0
	maxSum := nums[0]

	for _, n := range nums {
		if currSum+n > n {
			currSum += n
		} else {
			currSum = n
		}

		if currSum > maxSum {
			maxSum = currSum
		}
	}

	return maxSum
}
