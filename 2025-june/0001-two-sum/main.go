package main

import "fmt"

func main() {

	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// You can return the answer in any order.

	// CATEGORY: Arrays & Hashing
	// TIME COMPLEXITY: O(n)
	// SPACE COMPLEXITY: O(n)

	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)

	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {

	// start by traversing the array
	// 1. find target - currValue = ans
	// 2. check to see if ans is present in the map
	// 3. if not, add currValue and its index to the map
	// 4. if so, return the currvalue index and the target

	numsMap := make(map[int]int, len(nums))

	for currIndex, currValue := range nums {

		if otherIndex, isPresent := numsMap[target-currValue]; isPresent {
			// if it is present return both indices
			return []int{currIndex, otherIndex}
		}

		numsMap[currValue] = currIndex
	}

	return []int{}

}
