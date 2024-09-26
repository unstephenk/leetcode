package main

import "fmt"

func main() {
	// Given a string s, find the length of the longest substring without repeating characters.

	// s := "abcabcbb"
	s := "bbbbb"

	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	// start a for loop
	// check if char is present in hashmap and the index in the hashset is great than the left pointer
	// if true move the left pointer over one
	// mem := 0(n), space := 0(n)

	left, res := 0, 0

	hashSet := map[byte]int{}

	for i := 0; i < len(s); i++ {
		sChar := s[i]

		if _, isPresent := hashSet[sChar]; isPresent && hashSet[sChar] >= left {
			left = hashSet[sChar] + 1
		}

		res = max(res, i-left+1)
		hashSet[sChar] = i

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
