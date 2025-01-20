package main

import "fmt"

func main() {
	// There is an integer array nums sorted in ascending order (with distinct values).
	// Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is
	// [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and
	// become [4,5,6,7,0,1,2].
	// Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.
	// You must write an algorithm with O(log n) runtime complexity.

	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	res := search(nums, target)
	fmt.Println(res)

}

func search(nums []int, target int) int {
	// binary search
	// to increase search time, determine if the midpoint is in the left side or the right side
	// of the sorted array

	left, right := 0, len(nums)-1

	for left <= right {
		midpoint := (left + right) / 2
		if target == nums[midpoint] { // you have found the answer
			return midpoint
		}

		// determine if the midpoint is in the left/right side of the array
		if nums[left] <= nums[midpoint] { // in the left side
			if target > nums[midpoint] || target < nums[left] {
				left = midpoint + 1
			} else {
				right = midpoint - 1
			}
		} else { // in the right side
			if target < nums[midpoint] || target > nums[right] {
				right = midpoint - 1
			} else {
				left = midpoint + 1
			}
		}
	}

	return -1

}
