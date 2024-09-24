package main

import "fmt"

func main() {
	// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
	// You must write an algorithm that runs in O(n) time.

	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}

func longestConsecutive(nums []int) int {

	// create a set to look for left neighbors
	// loop through nums
	// if there is no neighbor than we have the start of a set
	// add 1 and update the longest var with the length

	neighborSet := make(map[int]int, len(nums))
	answer := 0

	for _, currValue := range nums {
		neighborSet[currValue]++
	}

	for _, numsValue := range nums {

		if _, isPresent := neighborSet[numsValue-1]; !isPresent {
			// {100, 4, 200, 1, 3, 2}
			currentNum := numsValue
			length := 0

			for {
				if _, isPresent := neighborSet[currentNum]; isPresent {
					length++
					currentNum++
				} else {
					break
				}
			}
			answer = max(answer, length)
		}
	}

	return answer
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
