package main

import "fmt"

func main() {
	// You are given a string s and an integer k. You can choose any character of the string and change it to any
	// other uppercase English character. You can perform this operation at most k times.
	// Return the length of the longest substring containing the same letter you can get after performing the above operations.
	s := "ABAB"
	k := 2

	res := characterReplacement(s, k)
	fmt.Println(res)
}

func characterReplacement(s string, k int) int {

}
