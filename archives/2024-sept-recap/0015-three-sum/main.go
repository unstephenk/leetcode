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
	// sort the array
	// start an outer loop, do not include the last character
	// start an inner loop, start at i + 1 and compare against the target

	var res [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // check for dups
			continue
		}

		target, left, right := -nums[i], 0, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target { // we have a match
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] { // check for dups again
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
