package main

import (
	"fmt"
	"sort"
)

func main() {

	// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that
	// i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
	// Notice that the solution set must not contain duplicate triplets.

	nums := []int{-1, 0, 1, 2, -1, -4}

	res := threeSum(nums)

	fmt.Println(res)

}

func threeSum(nums []int) [][]int {

	var res [][]int

	// sort the array
	sort.Ints(nums)

	// outer array, step through the nums
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // skip iteration if its a duplicates
			continue
		}

		target, left, right := -nums[i], i+1, len(nums)-1

		// inner loop
		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] { // skip duplicates
					left++
				}

			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}

	return res
}
