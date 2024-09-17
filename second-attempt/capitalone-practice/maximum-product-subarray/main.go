package main

import "fmt"

func main() {
	// Given an integer array, the task is to find the maximum product of any subarray.

	arr := []int{-2, 6, -3, -10, 0, 2}
	res := maxProduct(arr)
	fmt.Println(res)

}

func maxProduct(arr []int) int {
	// initate currMax, currMin, maxProd
	// step through the loop
	// assign currMax, currMin, maxProd
	// return maxProd

	currMax, currMin, maxProd := arr[0], arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		temp := max(currMax*arr[i], currMin*arr[i], arr[i])
		currMin = min(currMax*arr[i], currMin*arr[i], arr[i])
		currMax = temp
		maxProd = max(maxProd, currMax, maxProd)
	}

	return maxProd
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
