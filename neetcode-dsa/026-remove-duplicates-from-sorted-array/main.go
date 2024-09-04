package main

import "fmt"

func main() {
	// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
	// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
	// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
	// Return k.

	nums := []int{1, 1, 2}
	res := removeDuplicates(nums)
	fmt.Println(res)

}

func removeDuplicates(nums []int) int {
	// step through each element in the array
	// grab the ith and compare until the element changes
	// replace the same elements with 0
	// increment the answer
	// step to the next unique element

	left := 0

	for right := 0; right < len(nums); right++ {
		if nums[left] != nums[right] {
			nums[left+1] = nums[right]
			left++
		}
	}

	return left
}
