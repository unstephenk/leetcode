package main

import "fmt"

func main() {
	// Given an integer array nums, return true if any value appears at least twice in the array,
	// and return false if every element is distinct.

	nums := []int{1, 2, 3, 1}

	res := containsDuplicate(nums)

	fmt.Println(res)

}

func containsDuplicate(nums []int) bool {
	// create a map
	// iterate through the nums
	// check to see if it exists first
	// if it exists, then return false
	// if not, add the nums as a key to the map
	// return true

	numsMap := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, isPresent := numsMap[num]; isPresent {
			return true
		}

		numsMap[num] = true
	}

	return false

}
