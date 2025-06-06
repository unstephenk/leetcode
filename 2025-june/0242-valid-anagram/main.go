package main

import "fmt"

func main() {
	// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	// CATEGORY: Arrays & Hashing
	// TIME COMPLEXITY: O(2n) of O(n)

	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	// traverse the array of s and add each char to a map with occurences
	// traverse the array of t and remove from the map
	// if the char does not exist return true

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))

	for _, sChar := range s {
		sMap[sChar]++
	}

	for _, tChar := range t {
		if isPresent, _ := sMap[tChar]; isPresent == 0 {
			return false
		}
		sMap[tChar]--
	}

	return true
}
