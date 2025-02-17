package main

import "fmt"

func main() {
	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// You can return the answer in any order.

	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {
	// add all nums to a map and save their index
	// step through each element in the array
	// look for a num = target - element in the map
	// if found return the indices

	twoSumMap := make(map[int]int, len(nums))

	for numIndex, numVal := range nums {

		if otherIndex, isPresent := twoSumMap[target-numVal]; isPresent {
			return []int{otherIndex, numIndex}
		}

		twoSumMap[numVal] = numIndex
	}

	return []int{}
}
