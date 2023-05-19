package main

func maxProfit(prices []int) int {
	leftProfit := 0 //max profit
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice > prices[i-1] {
			minPrice = prices[i-1]
		}
		if leftProfit < prices[i]-minPrice {
			leftProfit = prices[i] - minPrice
		}
	}
	rightProfit := 0
	maxSellPrice := prices[len(prices)-1]
	for j := len(prices) - 2; j >= 0; j-- {
		if maxSellPrice < prices[j+1] {
			maxSellPrice = prices[j+1]
		}
		if rightProfit < maxSellPrice-prices[j] {
			rightProfit = maxSellPrice - prices[j]
		}
	}
	if leftProfit < rightProfit {
		return rightProfit
	}
	return leftProfit
}
