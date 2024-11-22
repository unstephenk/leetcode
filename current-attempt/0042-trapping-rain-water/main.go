package main

import "fmt"

func main() {
	// Given n non-negative integers representing an elevation map where the width of each bar is 1,
	// compute how much water it can trap after raining.

	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	res1 := trap(height)
	res2 := rainWaterPreCalc(height)
	fmt.Println(res1, res2)
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

func rainWaterPreCalc(height []int) int {
	// precalculate the highest columns for left and right
	// add them to arrays
	// step through the array again
	// calculate the volume

	res := 0
	lengthArr := len(height)

	// precalulate the left side
	left := make([]int, len(height))
	left[0] = height[0]
	// do not include the first element
	for i := 1; i < len(height); i++ {
		left[i] = max(left[i-1], height[i])
	}

	// do not include the last element
	right := make([]int, len(height))
	right[lengthArr-1] = height[lengthArr-1]
	for i := len(height) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	// both arrrays are populated, run calculation
	for i := 0; i < lengthArr; i++ {
		res += min(left[i], right[i]) - height[i]
	}

	return res

}
