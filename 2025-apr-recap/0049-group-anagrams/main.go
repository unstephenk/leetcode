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

}
