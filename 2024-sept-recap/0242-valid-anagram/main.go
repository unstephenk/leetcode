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
	if len(s) != len(t) {
		return false
	}

	sCharMap := make(map[rune]int, len(s))

	for _, letter := range s {
		sCharMap[letter]++
	}

	for _, letter := range t {
		sCharMap[letter]--

		if sCharMap[letter] < 0 {
			return false
		}
	}

	return true
}
