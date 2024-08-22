package main

import "fmt"

func main() {

	// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
	// Add keys to a hashmap and set the values as true. If true is already there, stop the process and return true

	// nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	nums := []int{1, 2, 3, 4}

	result := containsDuplicate(nums)

	fmt.Println(result)

}

func containsDuplicate(nums []int) bool {

	stephensMap := make(map[int]bool, len(nums))

	for _, item := range nums {
		if _, result := stephensMap[item]; result {
			return true
		}

		stephensMap[item] = true
	}

	return false
}
