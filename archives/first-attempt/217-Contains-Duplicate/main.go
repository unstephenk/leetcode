package main

import "fmt"

func main() {

	testArray := [5]int{1, 2, 3, 1}

	value := containsDuplicate(testArray)

	fmt.Println("The result was: ", value)
}

func containsDuplicate(nums [5]int) bool {

	stephensMap := make(map[int]bool, len(nums))

	for _, item := range nums {
		// check if this item alreadys exists in the map
		if _, ok := stephensMap[item]; ok {
			return true
		}

		stephensMap[item] = true
	}

	return false
}
