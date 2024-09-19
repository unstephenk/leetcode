package main

import "fmt"

func main() {
	// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

	nums := []int{3, 3}
	res := containsDuplicate(nums)
	fmt.Println(res)

}

func containsDuplicate(nums []int) bool {

	// traverse the array
	// before adding to the map, return false if this key already exists

	dupMap := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, isPresent := dupMap[num]; isPresent {
			return true
		}

		dupMap[num] = 1
	}

	return false
}
