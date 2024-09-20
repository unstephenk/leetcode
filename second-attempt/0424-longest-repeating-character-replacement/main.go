package main

import "fmt"

func main() {
	// You are given a string s and an integer k. You can choose any character of the string and
	// change it to any other uppercase English character. You can perform this operation at most k times.
	// Return the length of the longest substring containing the same letter you can get after performing the above operations.

	s, k := "ABAB", 2

	res := characterReplacement(s, k)

	fmt.Println(res)
}

func characterReplacement(s string, k int) int {

	// make a map of counts
	cnt := make(map[byte]int)

	res, left, maxf := 0, 0, 0

	for right := 0; right < len(s); right++ {
		// add to the map
		// update maxf
		// check if sliding window needs to move
		// update res
		// return res

		cnt[s[right]]++

		maxf = max(maxf, cnt[s[right]])

		if (right-left+1)-maxf > k {
			cnt[s[left]]--
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
