package main

import "fmt"

func main() {
	// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
}
