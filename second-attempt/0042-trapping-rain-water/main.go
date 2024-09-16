package main

import "fmt"

func main() {
	// Given n non-negative integers representing an elevation map where the width of each bar is 1,
	// compute how much water it can trap after raining.

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res := trap(height)
	fmt.Println(res)
}

func trap(height []int) int {
	// loop through the array
	// find the left side min
	// find the right side min
	// calulate the res
	// output the res

	res := 0

	// leave out the first and last element because water needs both sides to be retained
	for i := 1; i < len(height)-1; i++ {
		// find the min of the left side of the vessel
		left := height[i]
		for j := 0; j < i; j++ {
			left = max(left, height[j])
		}

		right := height[i]
		for j := i + 1; j < len(height); j++ {
			right = max(right, height[j])
		}

		res += min(left, right) - height[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
