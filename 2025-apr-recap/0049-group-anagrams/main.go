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

	masterMap := map[[26]int][]string{}

	for _, strsVal := range strs {
		tempArray := [26]int{}

		for _, strsChar := range strsVal {
			tempArray[strsChar-'a']++
		}

		masterMap[tempArray] = append(masterMap[tempArray], strsVal)
	}

	// the answers requires a different mapping
	answerArray := [][]string{}
	for _, masterMapVal := range masterMap {
		answerArray = append(answerArray, masterMapVal)
	}

	return answerArray

}
