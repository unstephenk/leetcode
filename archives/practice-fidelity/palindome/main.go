package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	inputString := "A man, a plan, a canal: Panama"

	res := isPalindrome(inputString)

	fmt.Println(res)

}

func isPalindrome(inputString string) bool {
	cleanUpString := func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	strConv := strings.Map(cleanUpString, inputString)

	i, j := 0, len(strConv)-1

	for i < j {
		if strConv[i] != strConv[j] {
			return false
		}
		i++
		j--
	}

	return true
}
