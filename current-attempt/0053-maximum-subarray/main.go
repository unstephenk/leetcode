package main

import "fmt"

func main() {

	// Given an integer array nums, find the subarray with the largest sum, and return its sum.

	testArray := []int{2, 3, -8, 7, -1, 2, 3}
	res := maxSubArray(testArray)

	fmt.Println(res)

}

func maxSubArray(nums []int) int {
	res := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum+nums[i], nums[i])
		res = max(currSum, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
