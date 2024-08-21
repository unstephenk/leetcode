package main

import "fmt"

func main() {

	// value := isAnagram("anagram", "nagaram")
	value := isAnagram("rat", "cat")

	fmt.Println("The result was: ", value)

}

func isAnagram(s string, t string) bool {
	// compare the lengths of the strings
	if len(s) != len(t) {
		return false
	}

	// add the chars of s to a map
	sMap := make(map[rune]int, len(s))

	// add s chars to the map
	for _, sChar := range s {
		sMap[sChar]++
	}

	for _, tChar := range t {
		sMap[tChar]--

		if sMap[tChar] < 0 {
			return false
		}
	}

	return true

}
