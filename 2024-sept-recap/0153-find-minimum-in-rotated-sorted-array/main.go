package main

import "fmt"

func main() {
	// Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
	// [4,5,6,7,0,1,2] if it was rotated 4 times.
	// [0,1,2,4,5,6,7] if it was rotated 7 times.
	// Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].
	// Given the sorted rotated array nums of unique elements, return the minimum element of this array.
	// You must write an algorithm that runs in O(log n) time.

	nums := []int{3, 4, 5, 1, 2}

	res := findMin(nums)
	fmt.Println(res)
}

func findMin(nums []int) int {

	// Modified Binary Search

	res := nums[0]
	l, r := 0, len(nums)-1

	for l <= r {
		// check to make see if the array is already sorted
		if nums[l] <= nums[r] {
			res = min(res, nums[l])
			break
		}

		// set the midpoint
		m := (l + r) / 2
		res = min(res, nums[m])

		// there are two portions of the array, the smaller portion and the larger portion
		// if the midpoint >= to the value at the beginning of the array, then we need to search the right portion of the array for smallest element
		if nums[m] >= nums[l] {
			// move the left pointer to the right of the midpoint
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return res
}

func min(a, b int) int {
	if b < a {
		return b
	} else {
		return a
	}

}
