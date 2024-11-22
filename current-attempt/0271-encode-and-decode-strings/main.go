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
	// The encoded string is then sent over the network and is decoded back to the original list of strings

	codec := Constructor()

	dummy_input := []string{"Hello", "World"}

	encodedString := codec.Encode(dummy_input)
	fmt.Println(encodedString)
	decodedString := codec.Decode(encodedString)
	fmt.Println(decodedString)
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {

	var res = strings.Builder{}

	for _, strChar := range strs {
		res.WriteString(strconv.Itoa(len(strChar))) // add the length to the beginning
		res.WriteString("#")
		res.WriteString(strChar)
	}

	return res.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {

	// 5#Hello#5World
	res := []string{}

	for len(strs) != 0 {

		letterCount := 0

		for _, strChar := range strs {
			if strChar == '#' {
				break
			}
			letterCount++
		}

		strLength, _ := strconv.Atoi(strs[:letterCount])

		res = append(res, strs[letterCount+1:letterCount+1+strLength])
		strs = strs[letterCount+1+strLength:]
	}

	return res

}
