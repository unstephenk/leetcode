package main

import "fmt"

func main() {
	// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.

	nums := []int{3, 2, 2, 3}
	val := 3

	res := removeElement(nums, val)
	fmt.Println(res)
}

func removeElement(nums []int, val int) int {

	k := 0

	for _, element := range nums {
		if element != val { // if the current element is not equal to k
			nums[k] = element
			k++
		}
		// if the element is equal to the k do nothing
	}

	return k

}
