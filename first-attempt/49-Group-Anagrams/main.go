package main

import "fmt"

func main() {

	value := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	// value := []string{""}
	// value := []string{"a"}

	groupAnagrams(value)

}

func groupAnagrams(strs []string) [][]string {

	// ["eat","tea","tan","ate","nat","bat"]

	// create a map, use the 26 ints as the value, use the string as the value
	// loop through the strings
	// for each string, add the values to a tempArray
	// add the key, value pair to the map
	// return the result as an array

	stephensMap := map[[26]int][]string{}

	for _, elementString := range strs {

		tempArray := [26]int{}

		for _, character := range elementString {
			tempArray[character-'a']++
		}

		stephensMap[tempArray] = append(stephensMap[tempArray], elementString)
	}

	res := [][]string{}

	for _, value := range stephensMap {
		res = append(res, value)
	}

	fmt.Println(res)

	return res

}
