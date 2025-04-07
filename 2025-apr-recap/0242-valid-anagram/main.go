package main

import "fmt"

func main() {
	// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	// Anagram = a word or phrase formed by rearranging the letters of another word or phrase,
	// typically using all the original letters exactly once. For example, "listen" and
	// "silent" are anagrams of each other.

	// CATEGORY: Arrays & Hashing

	s := "anagram"
	t := "nagaram"

	res := isAnagram(s, t)
	fmt.Println(res)
}

func isAnagram(s string, t string) bool {
	// add the letters in s to a map, use char as the key and occurences as the value
	// step through s and verify all the letter exist in the map, decrement after checking
	// if value is 0 and needs to be decremented, exit the lopp and return false

	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))

	for _, char := range s {
		sMap[char]++
	}

	for _, char := range t {
		sMap[char]--

		if sMap[char] < 0 {
			return false
		}
	}

	return true
}
