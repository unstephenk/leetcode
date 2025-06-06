package main

import "fmt"

func main() {
	// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	res := groupAnagrams(strs)
	fmt.Println(res)
}

func groupAnagrams(strs []string) [][]string {
	// create a array of 26 values starting with zeros
	// populate each position in the array with a value
	// set the new array as the key in map
	// traverse through the array, create a temp map and see if it exists in the overall map
	// { [0,0,0,0]: 'abc' }
	// [['eat', 'tea'], [tan]]

	res := map[[26]int][]string{}
	fmt.Println(res)

	for _, fullString := range strs {

		temp := [26]int{}

		for _, strsChar := range fullString {
			temp[strsChar-'a']++
		}

		res[temp] = append(res[temp], fullString)
	}
	fmt.Println(res)

	answer := [][]string{}

	for _, ana := range res {
		answer = append(answer, ana)
	}

	return answer

}
