package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
	// removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
	// Given a string s, return true if it is a palindrome, or false otherwise.

	s := "A man, a plan, a canal: Panama"
	// s := "race a car"

	res := isPalindrome(s)

	fmt.Println(res)
}

func isPalindrome(s string) bool {
	alphaNum := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	s = strings.Map(alphaNum, s)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true

}
