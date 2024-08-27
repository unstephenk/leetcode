package main

import "fmt"

func main() {

	// You are given an array prices where prices[i] is the price of a given stock on the ith day.
	// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
	// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)
	fmt.Println(res)

}

func maxProfit(prices []int) int {

	profit := 0
	cheapest := prices[0]

	for _, price := range prices {
		if price < cheapest {
			cheapest = price
		} else if (price - cheapest) > profit {
			profit = price - cheapest
		}
	}

	return profit
}
