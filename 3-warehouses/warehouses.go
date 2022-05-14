package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cnt int
	fmt.Scan(&cnt)

	var input string
	prices := make([]int, cnt)

	for i := 0; i < cnt; i++ {
		fmt.Scan(&input)
		prices[i], _ = strconv.Atoi(input)
	}

	dict := make(map[int]int)
	for i, price := range prices {
		maxPrice := price
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > maxPrice {
				maxPrice = prices[j]
			}
		}
		dict[i] = maxPrice
	}

	allPrice := 0
	for _, dayPrice := range dict {
		allPrice = allPrice + dayPrice
	}

	fmt.Println(allPrice)
}
