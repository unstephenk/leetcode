package main

import "fmt"

func main() {
	// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
	// You must write an algorithm that runs in O(n) time.

	// CATEGORY: Arrays & Hashing
	// TIME: O(n)
	// SPACE: O(n)

	nums := []int{100, 4, 200, 1, 3, 2}

	res := longestConsecutive(nums)

	fmt.Println(res)

}

func longestConsecutive(nums []int) int {
	numsMap := make(map[int]bool, len(nums))
	res := 0

	for _, numsValue := range nums {
		numsMap[numsValue] = true
	}

	for item := range numsMap {
		if _, ok := numsMap[item-1]; !ok { // this is the start of the sequence

			cnt := 1
			for _, ok := numsMap[item+cnt]; ok; _, ok = numsMap[item+cnt] {
				cnt++
			}

			res = max(cnt, res)
		}
	}

	return res
}

func max(newLength, longest int) int {
	if newLength > longest {
		return newLength
	}
	return longest
}
