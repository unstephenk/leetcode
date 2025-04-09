package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Constructor() Codec {
	return Codec{}
}

type Codec struct {
}

func main() {

	// Design an algorithm to encode a list of strings to a string
	// The encoded string is then sent over the network and is decoded back
	// to the original list of strings

	// CATEGORY: Arrays & Hashing
	// Add a delimeter such as the # sign
	// Add the total num of chars after the delimeter
	// Use For loop with Break

	codec := Constructor()

	dummy_input := []string{"Hello", "World"}

	encodedString := codec.Encode(dummy_input)
	fmt.Println(encodedString)
	decodedString := codec.Decode(encodedString)
	fmt.Println(decodedString)
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {

	res := strings.Builder{}

	for _, str := range strs {
		res.WriteString(str)
		res.WriteString("#")
		res.WriteString(strconv.Itoa(len(str)))
	}

	return res.String()

}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	// walk through the string until you find the delimeter
	// capture the value
	// grab all the chars starting from the beginning using the value
	// add the string to the answer array

	answer := []string{}

	for len(strs)
}
