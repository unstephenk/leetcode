package main

import "fmt"

func main() {

	// Trapping Rainwater Problem states that given an array of N non-negative integers arr[] representing an elevation map where the width of each bar is 1, compute how much water it can trap after rain.
	arr := []int{3, 0, 1, 0, 4, 0, 2}
	res1 := rainWaterNaive(arr)
	res2 := rainWaterPreCalc(arr)

	fmt.Println(res1)
	fmt.Println(res2)

}

func rainWaterNaive(arr []int) int {
	// loop through the array
	// find the left side min
	// find the right side min
	// calulate the res
	// output the res

	res := 0

	// leave out the first and last element because water needs both sides to be retained
	for i := 1; i < len(arr)-1; i++ {
		// find the min of the left side of the vessel
		left := arr[i]
		for j := 0; j < i; j++ {
			left = max(left, arr[j])
		}

		right := arr[i]
		for j := i + 1; j < len(arr); j++ {
			right = max(right, arr[j])
		}

		res += min(left, right) - arr[i]
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
	// precalculate the heights to the left and the right
	// calulate the volume
	n := len(height)

	left := make([]int, n)
	right := make([]int, n)
	res := 0

	//left side, do not include the first element
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], height[i])
	}

	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = max(right[i+1], height[i])
	}

	// calculation
	for i := 0; i < n; i++ {
		res += min(left[i], right[i]) - height[i]
	}

	return res

}
