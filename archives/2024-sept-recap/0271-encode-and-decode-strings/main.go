package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// Design an algorithm to encode a list of strings to a string.
	// The encoded string is then sent over the network and is decoded back to the original list of strings.

	// first append the strings together and add a special char that I have identified
	// next decode the string by looking for that special char and return the array of strings decoded

	// initalize the codec
	codec := Constructor()
	dummy_input := []string{"Hello", "World"}

	encoded := codec.Encode(dummy_input)

	fmt.Println("encoded: ", encoded)

	decoded := codec.Decode(encoded)

	fmt.Println("decoded: ", decoded)

}

type Codec struct {
	strs string
}

func Constructor() Codec {
	return Codec{}
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	// add a length of the string
	// add a special character

	res := strings.Builder{}

	for _, strsValue := range strs {
		res.WriteString(strconv.Itoa(len(strsValue)))
		res.WriteString("#")
		res.WriteString(strsValue)
	}

	return res.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {

	// first determine the lenght of the string
	// start collect the chars after the special char
	// add to the array

	ans := []string{}

	for len(strs) != 0 {

		charsInStringCount := 0

		for _, strsChar := range strs {
			// find the total num of ints to identify the string length
			if strsChar == '#' {
				break
			}
			charsInStringCount++
		}

		// remove the chars from the strs and add those to the answer array
		stringCount, _ := strconv.Atoi(strs[:charsInStringCount])
		fmt.Println(stringCount)

		ans = append(ans, strs[charsInStringCount+1:charsInStringCount+1+stringCount])
		strs = strs[charsInStringCount+1+stringCount:]
	}

	return ans

}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));
