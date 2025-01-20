package main

import "fmt"

func main() {
	// Given an array arr[], the task is to find the subarray that has the maximum sum and return its sum.

	testArray := []int{2, 3, -8, 7, -1, 2, 3}
	res := findSubArray(testArray)

	fmt.Println(res)
}

func findSubArray(arr []int) int {
	// kadane's algorithm

	// step through each item in the array
	// compare the maxsum to the current item + the maxsum
	// update the res if its higher
	// return the res

	res := arr[0]
	maxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		fmt.Println(arr[i], maxSum+arr[i], res)
		maxSum = max(maxSum+arr[i], arr[i])
		res = max(maxSum, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
