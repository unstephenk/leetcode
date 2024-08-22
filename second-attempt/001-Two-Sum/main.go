package main

import "fmt"

func main() {

	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// You can return the answer in any order.

	// nums := []int{2, 7, 11, 15}
	// target := 9

	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {

	stephensMap := make(map[int]int, len(nums))

	// add the values to the hashmap as a key
	for currIndex, currValue := range nums {

		// check to see if the value exists in the hashmap of target - currValue
		if otherIndex, isPresent := stephensMap[target-currValue]; isPresent {
			// if it does than return the indices of both values
			return []int{otherIndex, currIndex}
		}

		// add the nums value to the hashmap as the key with index
		stephensMap[currValue] = currIndex
	}

	return []int{}
}
