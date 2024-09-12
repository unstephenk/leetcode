package main

import (
	"fmt"
	"strings"
)

func main() {
	// Implement a function to justify a list of words within a given width.
	// In word processing, justify refers to the alignment of text such that
	// it is evenly distributed between the left and right margins of a page.
	// When text is justified, extra spacing is added between words to ensure
	// that both edges of the text line up evenly, creating a clean and professional appearance.
	width := 32
	listOfWords := []string{"the", "quick", "fox", "jumps"} // 16 characters

	res := justifyWords(width, listOfWords)
	fmt.Println(res)

}

func justifyWords(width int, listOfWords []string) string {
	countOfChars := 0
	requiredSpaces := len(listOfWords) - 1
	res := strings.Builder{}

	for _, s := range listOfWords {
		countOfChars += len(s)
	}

	// add the chars and the requrired spaces and then subtract the width to return the num of additional spaces needed
	additionalSpaces := (width - countOfChars) / (len(listOfWords) - 1)

	fmt.Println(width, countOfChars, requiredSpaces, (len(listOfWords) - 1), additionalSpaces)

	for i := 0; i <= len(listOfWords)-1; i++ {
		res.WriteString(listOfWords[i])

		if i == (len(listOfWords) - 1) { // at the last word
			continue
		} else {
			for i := 1; i <= additionalSpaces; i++ {
				res.WriteString("_")
			}
		}

	}

	fmt.Println(len(res.String()))
	return res.String()
}
