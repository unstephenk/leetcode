package main

import "fmt"

func main() {
	// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
	// [4,5,6,7,0,1,2] if it was rotated 4 times.
	// [0,1,2,4,5,6,7] if it was rotated 7 times.
	// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
	// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
	// You must write an algorithm that runs in O(log n) time.

	// nums := []int{3, 4, 5, 1, 2}
	nums := []int{2, 3, 4, 5, 1}
	res := findMin(nums)
	fmt.Println(res)
}

func findMin(nums []int) int {

	// This is a BST algorithm
	// start at the middle point
	// determine if the middle point is part of the left sorted array or part of the right sorted array
	// if its part of the left, move the left pointer to the next element and search for the lowest on that side
	// if its part of the right, move the left pointer to the beginning of the array and start searching there for the lowest element

	// first assign the left pointer at the beginning and the right pointer at the end
	// find the middle pointer by / 2 and assign that as the res
	// determine if you are in the left or the right by checking nums[m] >= nums[leftPointer]

	left, right := 0, len(nums)-1

	for left < right {
		midpoint := (left + right) / 2

		if nums[midpoint] > nums[right] {
			left = midpoint + 1
		} else {
			right = midpoint
		}

	}
	return nums[left]
}
