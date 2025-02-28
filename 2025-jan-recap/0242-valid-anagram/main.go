package main

import "fmt"

func main() {
	// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {

	// push all the chars from string S to a map
	// walk through string T and remove the element from the map
	// of the element value is less than 1, return false

	sMap := make(map[rune]int, len(s))

	// first check to make sure the lens match
	if len(s) != len(t) {
		return false
	}

	for _, sChar := range s {
		sMap[sChar]++
	}

	for _, tChar := range t {
		if count, isPresent := sMap[tChar]; isPresent && count > 0 {
			sMap[tChar]--
		} else {
			return false
		}
	}

	return true
}
