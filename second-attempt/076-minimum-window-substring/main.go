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

	// create a map of t string (need)
	// step through the string s
	// compare what we have for that char vs what we need
	// if equal, update the have map
	// then compare the have total to the need total and if equal compare the res to check if it needs an update
	// if not equal, do not update the have map and add a new char to the window

	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	needMap := make([]int, 128)
	needMapCount := len(t)

	start, end := 0, 0
	minLen, startIndex := int(^uint(0)>>1), 0

	for _, char := range t { // add the chars to the needmap
		needMap[char]++
	}

	for end < len(s) {

		// ADOBECODEBANC
		if needMap[s[end]] > 0 { // decrement the needMapCount if the char is present
			needMapCount--
		}

		needMap[s[end]]-- // decrement the count from the needMap
		end++             // move the end pointer right

		for needMapCount == 0 {
			if end-start < minLen {
				startIndex = start
				minLen = end - start
			}

			if needMap[s[start]] == 0 { // increment the
				needMapCount++
			}
			needMap[s[start]]++
			start++
		}
	}

	if minLen == int(^uint(0)>>1) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
