package main

import "fmt"

func main() {

	// in Unicode, a-z represents number 97-122. Instead of creating an array up to 200 values, we know that in this problem we are only
	// given lowercase a-z which is 26 characters so we could simply subtract the current letter from a(97) to create an array with
	// 26 characters max improving performance of the code.

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	results := groupAnagrams(strs)

	fmt.Println(results)

}

func groupAnagrams(strs []string) [][]string {

	stephensMap := map[[26]int][]string{}

	for _, strsValue := range strs {

		tempArray := [26]int{}

		for _, charValue := range strsValue {
			tempArray[charValue-'a']++
		}

		stephensMap[tempArray] = append(stephensMap[tempArray], strsValue)
	}

	// answer requires that you return an array of an array of strings
	// do the conversion here
	res := [][]string{}

	for _, currValue := range stephensMap {
		res = append(res, currValue)
	}

	return res
}
