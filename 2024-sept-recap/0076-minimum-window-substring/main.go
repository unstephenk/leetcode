package main

import "fmt"

func main() {
	// Given two strings s and t of lengths m and n respectively, return the minimum window substring
	// of s such that every character in t (including duplicates) is included in the window. If there is no such
	// substring, return the empty string "".
	// The testcases will be generated such that the answer is unique.

	s := "ADOBECODEBANC"
	t := "ABC"

	res := minWindow(s, t)
	fmt.Println(res)
}

func minWindow(s string, t string) string {
	if t == "" { // handle edge case, if t is empty return empty
		return ""
	}

	// hashmap for window

	// hashmap to hold the chars in t
	tMap := map[rune]int{}

	for _, tChar := range t {
		tMap[tChar]++
	}

	have, need := 0, len(tMap) // len(tMap) returns the length of unique chars only

	for right := 0; right < len(s); right++ {
		currentChar := s[right]

	}
}
