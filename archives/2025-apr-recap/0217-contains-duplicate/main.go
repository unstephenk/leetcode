package main

import "fmt"

func main() {
	// Given an integer array nums, return true if any value appears at least twice in
	// the array, and return false if every element is distinct.

	// CATEGORY: Arrays & Hashing

	nums := []int{1, 2, 3, 1}
	res := containsDups(nums)
	fmt.Println(res)

}

func containsDups(nums []int) bool {
	// add it to the map and increment
	// check the value and if more than 2, return true

	numsMap := make(map[int]int, len(nums))

	for _, numsVal := range nums {
		if _, isPresent := numsMap[numsVal]; isPresent {
			return true
		}

		numsMap[numsVal] = 1
	}

	return false
}
