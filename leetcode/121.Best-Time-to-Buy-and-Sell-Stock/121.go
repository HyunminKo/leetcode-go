package leetcode

import "math"

func maxProfit(prices []int) int {
	max := 0
	min := math.MaxInt

	size := len(prices)

	for i := 0; i < size; i++ {
		if prices[i] < min {
			min = prices[i]
		}
		current := prices[i] - min
		if current >= max {
			max = current
		}
	}

	return max
}
