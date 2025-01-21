package main

import "fmt"

func main() {
	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// You can return the answer in any order.

	nums, target := []int{2, 7, 11, 15}, 9

	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {
	// step through the nums and add them to a map
	// check to see if the curr target minus the current value exists in the map
	// if not, continue
	// if so, add the curr num and the new num to the ans array

	numsMap := make(map[int]int, len(nums))

	for currIndex, currValue := range nums {
		if newIndex, isPresent := numsMap[target-currValue]; isPresent {
			return []int{currIndex, newIndex}
		}

		numsMap[currValue] = currIndex
	}

	return []int{}
}
