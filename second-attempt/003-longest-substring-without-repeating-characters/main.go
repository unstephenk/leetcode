package main

import (
	"fmt"
)

func main() {
	// Given a string s, find the length of the longest substring without repeating characters.
	// Sliding window

	s := "abcabcbb"

	res := lengthOfLongestSubstring(s)

	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	start := 0
	longest := 0

	set := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sChar := s[i]

		// check for dups, move start point
		if _, isPresent := set[sChar]; isPresent && set[sChar] >= start {
			start = set[sChar] + 1
		}

		longest = max(longest, i-start+1)
		set[sChar] = i
	}

	return longest
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
