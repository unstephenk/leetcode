package main

import (
	"fmt"
)

func main() {
	// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

	s, t := "anagram", "nagaram"

	res := validAnagram(s, t)
	fmt.Println(res)

}

func validAnagram(s, t string) bool {
	// step through s, and add it to a map
	// step through t, and add it to the map
	// check to see if it already exists
	// if not, return false
	// return true

	mapOfS := make(map[rune]int, len(s))

	if len(s) != len(t) {
		return false
	}

	for _, letter := range s {
		mapOfS[rune(letter)]++
	}

	for _, letter := range t {
		mapOfS[rune(letter)]--

		if mapOfS[rune(letter)] < 0 {
			return false
		}
	}

	return true

}
