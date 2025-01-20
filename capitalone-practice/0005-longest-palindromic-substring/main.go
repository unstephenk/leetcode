package main

import (
	"fmt"
)

func main() {
	// Given a string s, return the longest palindromic substring in s.
	s := "babad"
	res := longestPalindrome(s)
	fmt.Println(res)
}

func longestPalindrome(s string) string {
	// 2 cases to consider, odd and even lengths

	res := ""
	str := ""
	for start := 0; start < len(s); start++ {
		for end := start; end < len(s); end++ {
			str = s[start : end+1]
			if Check(str) {
				if len(str) > len(res) {
					res = str
				}
			}
		}
	}
	return res
}

func Check(str string) bool {
	for start, end := 0, len(str)-1; start < end; start++ {
		if str[start] != str[end] {
			return false
		}
		end--
	}
	return true
}
