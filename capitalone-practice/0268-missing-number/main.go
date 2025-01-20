package main

import "fmt"

func main() {
	// Given an array nums containing n distinct numbers in
	// the range [0, n], return the only number in the range that is missing from the array.
	nums := []int{3, 0, 1}

	res := missingNumber(nums)

	fmt.Println(res)

}

func missingNumber(nums []int) int {
	// SUMMATION LAW
	// find the expected sum
	// find the sum of all the array elements
	// subtract expected sum from the sum of all of the elements

	expectedSum, totalSum := 0, 0
	distinctValues := len(nums)

	for _, num := range nums {
		totalSum += num
	}

	expectedSum = (distinctValues * (distinctValues + 1)) / 2

	return (expectedSum - totalSum)

}
