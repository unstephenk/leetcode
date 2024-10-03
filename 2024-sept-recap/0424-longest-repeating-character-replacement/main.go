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

	// step through the string
	// update the count in the map
	// update the max freq by looking at the freq in the map
	// if k >= maxf then update res
	// else move the left pointer over 1

	mapCount := make(map[byte]int, 26)

	left, maxf, res := 0, 0, 0

	for right := 0; right < len(s); right++ {
		mapCount[s[right]]++

		maxf = max(maxf, mapCount[s[right]])

		if (right-left+1)-maxf > k {
			mapCount[s[right]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
