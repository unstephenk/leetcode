package main

import "fmt"

func main() {
	// Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.
	// You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

	// CATEGORY: Arrays & Hashing
	// TIME COMPLEXITY:

	nums := []int{1, 2, 0}
	res := firstMissingPositive(nums)
	fmt.Println(res)

}

func firstMissingPositive(nums []int) int {

}
