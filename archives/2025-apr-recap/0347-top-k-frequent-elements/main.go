package main

import "fmt"

func main() {

	// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

	// CATEGORY: Arrays & Hashing
	// TIME: O(n) - n is the size of the input array
	// SPACE: O(n) - n is the size of the input array
	// ALGO: Bucket Sort

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	result := topKFrequent(nums, k)

	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {

	// create a map and use the nums value as the key
	// increment the value of the key everytime we find a dup
	// convert the first map to a Bucket Sort
	// start at the last element in the bucket sort and find the k most frequent values

	cntMap := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	for _, numsValue := range nums {
		cntMap[numsValue]++
	}

	for num, cnt := range cntMap {
		freq[cnt] = append(freq[cnt], num)
	}

	res := []int{}

	for i := len(freq) - 1; i > 0; i-- {
		for _, val := range freq[i] {
			res = append(res, val)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
