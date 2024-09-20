package main

import "fmt"

func main() {

	// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
	// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
	// You must write an algorithm that runs in O(n) time and without using the division operation.

	nums := []int{1, 2, 3, 4}
	res := productExceptSelf(nums)
	fmt.Println(res)

}

func productExceptSelf(nums []int) []int {

	answer := make([]int, len(nums))

	prefixPointer := 1

	for currIndex, currValue := range nums {
		answer[currIndex] = prefixPointer
		prefixPointer *= currValue
	}

	postfixPointer := 1

	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfixPointer
		postfixPointer *= nums[i]
	}

	return answer

}
