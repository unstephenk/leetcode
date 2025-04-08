package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	inputString := "this is a car race"

	res := reverseString(inputString)
	fmt.Println(inputString)
	fmt.Println(res)
}

func reverseString(inputString string) string {
	checkForInvalidChars := func(r rune) rune {
		// if its not a letter or number, return -1 (ignored)
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return -1
		}

		return unicode.ToLower(r)
	}

	// use String Map to map the characters to a new string
	strConv := strings.Map(checkForInvalidChars, inputString)
	fmt.Println(strConv)

	// step through the chars starting at the end, add to a new string
	i, j := 0, len(strConv)-1

	var ansString = strings.Builder{}

	for j >= i {
		ansString.WriteRune(rune(strConv[j]))
		j--
	}

	return ansString.String()
}
