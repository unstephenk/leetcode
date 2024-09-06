package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var nameOfFile string

	// request the name of the file
	fmt.Println("Enter the name of the file:")
	reader := bufio.NewReader(os.Stdin)
	nameOfFile, _ = reader.ReadString('\n')
	nameOfFile = strings.TrimSpace(nameOfFile) // Remove trailing newline

	res := convertToNewString(nameOfFile)

	fmt.Println(res)

	// step through the string and find the first period
	// convert the pervious numbers to 1000
	// add the numbers to a new string
	// next, everything after the first space, add until you find another space
	// convert to lowerCase
}

func convertToNewString(s string) string {
	left := 0

	var res = strings.Builder{}

	for right := 0; right < len(s); right++ {

		if s[right] == '.' {

			if left-(right-1) < 3 {
				addZeros := 3 - (left - (right - 1))
				for i := 1; i <= addZeros; i++ {
					res.WriteString("0")
				}
			}

			res.WriteString(s[left:right])
			left = right + 1
		}

		if s[right] == ' ' {
			res.WriteString("-")
		}

		if unicode.IsUpper(rune(s[right])) || unicode.IsLower(rune(s[right])) {

			temp := string(unicode.ToLower(rune(s[right])))
			res.WriteString(temp)
		}
	}

	return res.String()
}
