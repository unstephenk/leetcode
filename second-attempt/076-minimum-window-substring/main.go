package main

import "fmt"

func main() {
	// Given two strings s and t of lengths m and n respectively, return the minimum window substring
	// of s such that every character in t (including duplicates) is included in the window. If there is no such substring, return the empty string "".
	// The testcases will be generated such that the answer is unique.

	s := "ADOBECODEBANC"
	t := "ABC"

	res := minWindow(s, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	// add t to a map
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	mapS := make([]int, 128)
	count := len(t)
	start, end := 0, 0
	minLen, startIndex := int(^uint(0)>>1), 0
	/// UPVOTE !
	for _, char := range t {
		mapS[char]++
	}

	for end < len(s) {
		if mapS[s[end]] > 0 {
			count--
		}
		mapS[s[end]]--
		end++

		for count == 0 {
			if end-start < minLen {
				startIndex = start
				minLen = end - start
			}

			if mapS[s[start]] == 0 {
				count++
			}
			mapS[s[start]]++
			start++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
