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
	// sliding window
	// start the left pointer at 0 and right at 0+1
	// if right is less than the left, move the left and right pointer
	// else calculate and update the max profit

	left, right, maxProfit := 0, 1, 0

	for right <= len(prices)-1 {
		if prices[left] < prices[right] {
			maxProfit = max(maxProfit, prices[right]-prices[left])
		} else {
			left = right
		}
		right++

	}

	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
