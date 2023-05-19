package main

func maxProfit2(prices []int) int {
	return buySell(prices, 0, len(prices)-1)
}

func buySell(prices []int, i, j int) int {
	if i >= j {
		return 0
	} else if i+1 == j {
		d := prices[j] - prices[i]
		if d > 0 {
			return d
		} else {
			return 0
		}
	}

	n := int((j - i) / 2)
	left := buySell(prices, i, i+n)
	right := buySell(prices, i+n, j)
	return left + right
}
