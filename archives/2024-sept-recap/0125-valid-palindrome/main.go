package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
	// it reads the same forward and backward. Alphanumeric characters include letters and numbers.
	// Given a string s, return true if it is a palindrome, or false otherwise.
	s := "A man, a plan, a canal: Panama"
	res := isPalindrome(s)
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	// filter the string and remove anything that is not a rune or a number
	// start a pointer on the left and one on the right
	// compare and return true or false

	alphaNum := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	strConv := strings.Map(alphaNum, s)

	left := 0
	right := len(strConv) - 1

	for left < right {
		if strConv[left] != strConv[right] {
			return false
		}

		left++
		right--
	}

	return true
}
