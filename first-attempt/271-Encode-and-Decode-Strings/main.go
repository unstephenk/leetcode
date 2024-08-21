package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	value := []string{"Hello", "World"}

	c := Codec{}

	encoded := c.Encode(value)
	decoded := c.Decode(encoded)

	fmt.Println(decoded)
}

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	// deconstruct the array of strings
	// add the total number of elements per array + the delimeter which will be #

	// passed on to the next machine
	var encodedString = strings.Builder{}

	for _, value := range strs {
		encodedString.WriteString(strconv.Itoa(len(value)))
		encodedString.WriteString("#")
		encodedString.WriteString(value)
	}
	fmt.Println(encodedString.String())
	return encodedString.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {

	res := []string{}

	// capture edge case of array equal to 0
	for len(strs) != 0 {

		cnt := 0 // integers before the delimiter

		for _, char := range strs { // step through the string
			if char == rune('#') { // break out when it reaches the delimeter
				break
			}
			cnt++ // otherwise, increase the count
		}

		length, _ := strconv.Atoi(strs[:cnt])

		res = append(res, strs[cnt+1:cnt+1+length])
		strs = strs[cnt+1+length:] // remove the length from strs and start again
	}

	return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
