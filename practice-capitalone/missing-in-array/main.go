package main

import "fmt"

func main() {
	// Given an array arr[] of size N-1 with integers in the range of [1, N], the task is to find the missing number from the first N integers.
	// Note: There are no duplicates in the list.

	// testArray := []int{1, 2, 4, 6, 3, 7, 8}
	// endNum := 8

	testArray := []int{1, 2, 3, 5}
	endNum := 5

	res := missingNumber(testArray, endNum)
	fmt.Println(res)

}

func missingNumber(testArray []int, endNum int) int {
	// SUMMATION LAW
	// find the expected sum
	// find the sum of all the array elements
	// subtract expected sum from the sum of all of the elements

	expectedSum, totalSum := 0, 0

	for _, num := range testArray {
		totalSum += num
	}

	expectedSum = (endNum * (endNum + 1)) / 2

	return (expectedSum - totalSum)

}
