package main

func maxProfit(prices []int) int {
	profit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			profit = max(profit, prices[i]-minPrice)
		}
	}
	return profit
}
