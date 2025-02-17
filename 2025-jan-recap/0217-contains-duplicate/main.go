package main

import "fmt"

func main() {
	// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

	nums := []int{1, 2, 3, 1}
	res := containsDups(nums)
	fmt.Println(res)

}

func containsDups(nums []int) bool {
	// Add each to a map

	duplicateMap := make(map[int]bool, len(nums))

	for _, numsVal := range nums {
		// check to see if the value already exists
		if _, isPresent := duplicateMap[numsVal]; isPresent {
			return true
		}

		duplicateMap[numsVal] = true
	}

	return false

}
