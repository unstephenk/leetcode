package main

import "fmt"

func main() {
	// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
	// Find two lines that together with the x-axis form a container, such that the container contains the most water.
	// Return the maximum amount of water a container can store.
	// Notice that you may not slant the container.

	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	res := maxArea(height)

	fmt.Println(res)
}

func maxArea(height []int) int {
	// left, right := 0, len(height)-1 // set left and right pointers
	// var res int

	n := len(height)

	var left, res int
	right := n - 1

	for left < right {
		// find the minHeight
		// find the area => minHeight * total of remaining heights
		// update the answer if its greater
		area := minHeight(height[left], height[right]) * (right - left)

		if area > res {
			res = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}

	}

	return res
}

func minHeight(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
