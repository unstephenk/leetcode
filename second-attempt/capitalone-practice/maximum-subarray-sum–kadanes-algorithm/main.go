package main

import "fmt"

func main() {
	// Given an array arr[], the task is to find the subarray that has the maximum sum and return its sum.

	testArray := []int{2, 3, -8, 7, -1, 2, 3}
	res := findSubArray(testArray)

	fmt.Println(res)
}

func findSubArray(arr []int) int {
	res := arr[0]
	currSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currSum = max(currSum+arr[i], arr[i])
		res = max(currSum, res)
		fmt.Println(arr[i], currSum, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
