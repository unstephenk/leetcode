package main

import "fmt"

func main() {

	// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
	// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

	// for the first string, add each letter in the anagram to a hashmap using the letter as the key and the occurence as the value
	// run a loop through the second string and decrement the value
	// if the value is less than zero, return false

	// s := "rat"
	// t := "car"

	s := "anagram"
	t := "nagaram"

	result := isAnagram(s, t)

	fmt.Println(result)
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	stephenMap := make(map[rune]int, len(s))

	for _, sChar := range s {
		stephenMap[sChar]++
	}

	for _, tChar := range t {
		stephenMap[tChar]--

		if stephenMap[tChar] < 0 {
			return false
		}
	}

	return true

}
