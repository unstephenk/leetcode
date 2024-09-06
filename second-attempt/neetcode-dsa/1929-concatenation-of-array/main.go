package main

import "fmt"

func main() {
	// Given an integer array nums of length n, you want to create an array ans of length 2n where ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n (0-indexed).
	// Specifically, ans is the concatenation of two nums arrays.
	// Return the array ans.

	nums := []int{1, 2, 1}

	res := getConcatenation(nums)

	fmt.Println(res)
}

func getConcatenation(nums []int) []int {

	// create a new array
	// step through nums and add it to the new array twice

	ans := make([]int, len(nums)*2)

	for position, currNum := range nums {
		ans[position] = currNum
		ans[position+len(nums)] = currNum
	}

	return ans

}
