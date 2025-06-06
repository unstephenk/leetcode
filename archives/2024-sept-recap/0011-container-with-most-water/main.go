package main

import "fmt"

func main() {
	// You are given an integer array height of length n. There are n vertical lines drawn such that the two
	// endpoints of the ith line are (i, 0) and (i, height[i]).
	// Find two lines that together with the x-axis form a container, such that the container contains the most water.
	// Return the maximum amount of water a container can store.
	// Notice that you may not slant the container.
	// height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height := []int{2, 3, 4, 5, 18, 17, 6}

	res := maxArea(height)
	fmt.Println(res)
}

func maxArea(height []int) int {
	// create left pointer at 0 and right at the end
	// start a for loop and calculate the area, update the greatest area
	// find the minimum of left and right and shift one of the pointers

	res := 0

	left, right := 0, len(height)-1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		res = max(res, area)

		if height[left] < height[right] || height[left] == height[right] {
			left++
		} else if height[left] > height[right] {
			right--
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
