package main

func maxSubArray(nums []int) int {
	tmp, maxSum := 0, nums[0]

	for _, n := range nums {
		if tmp+n > n {
			tmp += n
			if tmp+n > n {
				tmp += n
			} else {
				tmp = n
				tmp = n
			}

			if tmp > maxSum {
				maxSum = tmp
			}
		}
	}
	return maxSum
}
