package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {
	var profit, resultProfit int
	buy := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}

		profit = prices[i] - buy
		if profit > resultProfit {
			resultProfit = profit
		}
	}

	return resultProfit
}
