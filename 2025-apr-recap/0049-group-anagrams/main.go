package main

import "fmt"

func main() {
	// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
	// Example Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	// CATEGORY: Arrays & Hashing

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {

	masterMap := map[[26]int][]string{}

	for _, strsWord := range strs {

		tempArray := [26]int{}

		for _, strsChar := range strsWord {
			tempArray[strsChar-'a']++
		}

		masterMap[tempArray] = append(masterMap[tempArray], strsWord)
	}

	// answer requires a different format
	answer := [][]string{}
	for _, masterMapVal := range masterMap {
		answer = append(answer, masterMapVal)
	}

	return answer
}
