package main

import "fmt"

func main() {
	// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

	// Example Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {

	// create a map using 26 0's
	// iterate through the string
	// create a temp map
	// add the chars to the temp map
	// add the temp map to the res at the 26 char key with the string as the value

	res := map[[26]int][]string{}

	for _, fullstring := range strs {

		tempArray := [26]int{}

		for _, strChar := range fullstring {
			tempArray[strChar-'a']++
		}

		res[tempArray] = append(res[tempArray], fullstring)
	}

	answer := [][]string{}

	for _, ana := range res {
		answer = append(answer, ana)
	}

	return answer
}
