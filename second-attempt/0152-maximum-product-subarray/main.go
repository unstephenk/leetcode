package main

import (
	"fmt"
)

func main() {
	// Given an integer array, the task is to find the maximum product of any subarray.

	arr := []int{-2, 6, -3, -10, 0, 2}
	res := maxProduct(arr)
	fmt.Println(res)

}

func maxProduct(nums []int) int {
	// initate currMax, currMin, maxProd
	// step through the loop
	// assign currMax, currMin, maxProd
	// return maxProd

	currMax, currMin, maxProd := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			currMax, currMin = currMin, currMax
		}

		currMax = max(currMax*nums[i], nums[i])
		currMin = min(currMin*nums[i], nums[i])
		maxProd = max(currMax, maxProd)
	}

	return maxProd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
