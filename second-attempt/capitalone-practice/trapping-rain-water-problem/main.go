package main

import "fmt"

func main() {

	// Trapping Rainwater Problem states that given an array of N non-negative integers arr[] representing an elevation map where the width of each bar is 1, compute how much water it can trap after rain.
	arr := []int{3, 0, 1, 0, 4, 0, 2}
	res := rainWater(arr)

	fmt.Println(res)

}

func rainWater(arr []int) int {
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
