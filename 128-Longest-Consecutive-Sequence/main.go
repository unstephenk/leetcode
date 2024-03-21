package main

import "fmt"

func main() {
	value := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(value)

	fmt.Println(res)
}

func longestConsecutive(nums []int) int {
	// convert the array to map to make it simpler to search
	// step through the map
	// if the item in the map-1 does not exist in the map than its the start of the sequence
	// keep searching the map for item + 1 and increment the counter
	// at the end, compare the existing length and update if its greater

	stephensMap := make(map[int]struct{}, len(nums))

	// converts the array to a map
	for _, value := range nums {
		stephensMap[value] = struct{}{}
	}

	var longest int

	// step through each item in the map
	for item := range stephensMap {
		// if the pervious number does not exist, then the number is considered the start of a sequence
		if _, ok := stephensMap[item-1]; !ok {
			// set the length = 1
			length := 1
			// does the map contain more numbers in the sequence
			for _, ok := stephensMap[item+length]; ok; _, ok = stephensMap[item+length] {
				length++
			}
			longest = max(longest, length)

		}
	}

	return longest
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
