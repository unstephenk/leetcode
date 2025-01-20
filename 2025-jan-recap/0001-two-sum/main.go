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
	// step through the map
	// check to see if the curr nums minus the target exists in the map
	// if not, continue
	// if so, add the curr num and the new num to the ans array

	numsMap := make(map[int]int, len(nums))
	ansArray := []int{}

	for _, num := range nums {
		numsMap[num]++
	}

	for _, num := range nums {
		if _, isPresent := numsMap[target-num]; isPresent {
			ansArray = append(ansArray, numsMap[target-num])
			ansArray = append(ansArray, num)
		}
	}

	return ansArray
}
