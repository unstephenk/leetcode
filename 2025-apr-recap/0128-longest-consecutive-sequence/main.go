package main

import "fmt"

func main() {
	// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
	// You must write an algorithm that runs in O(n) time.

	// CATEGORY: Arrays & Hashing

	nums := []int{100, 4, 200, 1, 3, 2}

	res := longestConsecutive(nums)

	fmt.Println(res)

}

func longestConsecutive(nums []int) int {
}
