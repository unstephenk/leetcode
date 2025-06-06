package main

import "fmt"

func main() {
	// Given an integer array nums, return true if any value appears at least twice in the array,
	// and return false if every element is distinct.

	// CATEGORY: Arrays & Hashing
	// TIME COMPLEXITY: O(n)

	nums := []int{1, 2, 3, 1}
	res := containsDuplicate(nums)
	fmt.Println(res)
}

func containsDuplicate(nums []int) bool {
	// traverse the array
	// if the key already exists in the map, return false

	numsMap := make(map[int]bool, len(nums))

	for _, numsValue := range nums {
		if _, isPresent := numsMap[numsValue]; isPresent {
			return true
		}

		numsMap[numsValue] = true
	}

	return false

}
