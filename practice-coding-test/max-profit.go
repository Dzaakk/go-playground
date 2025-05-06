package main

// best time to buy and sell stock

func maxProfit(prices []int) int {
	min := prices[0]
	maxProfit := 0

	for i, n := range prices {
		if i == 0 {
			continue
		}

		if min > n {
			min = n
			continue
		} else {
			tmp := n - min
			if maxProfit >= tmp {
				continue
			} else {
				maxProfit = tmp
			}
		}
	}
	return maxProfit
}

// sliding window
func maxProfit2(prices []int) int {
	left := 0
	right := 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			left = right
		}
		right++
	}

	return maxProfit
}
