package main

import "fmt"

func main() {
	// Given an array of strings strs, group the anagrams together.
	// You can return the answer in any order.

	// CATEGORY: Arrays & Hashing
	// TIME COMPLEXITY:

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams(strs)
	fmt.Println(res)

}

func groupAnagrams(strs []string) [][]string {
	// travese the array
	// traverse the string
	// for each rune in the string, convert to its alpha value in a 26 index array
	// add that array and the string value to the hashmap
	// convert the hashmap to a nested array and return the answer

	anagramMap := map[[26]int][]string{}

	for _, strArray := range strs {

		tempArray := [26]int{}

		for _, charVal := range strArray {

			tempArray[charVal-'a']++
		}

		anagramMap[tempArray] = append(anagramMap[tempArray], strArray)
	}

	ansArray := [][]string{}

	for _, mapStrings := range anagramMap {
		ansArray = append(ansArray, mapStrings)
	}

	return ansArray
}
