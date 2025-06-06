package main

import "fmt"

func main() {
	// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
	// You may assume that each input would have exactly one solution, and you may not use the same element twice.
	// You can return the answer in any order.

	// CATEGORY: Arrays & Hashing

	nums := []int{2, 7, 11, 15}
	target := 9

	res := twoSum(nums, target)
	fmt.Println(res)

}

func twoSum(nums []int, target int) []int {

	// create map of all the values using the value as the key
	// take the value at the current index and calculate: target - value = answer
	// if the answer exists in the map, return both values

	numsMap := make(map[int]int, len(nums))

	for numPos, numVal := range nums {
		if otherPos, isPresent := numsMap[target-numVal]; isPresent {
			return []int{otherPos, numPos}
		}
		numsMap[numVal] = numPos
	}

	return []int{}
}
