package main

import "fmt"

func main() {

	//nums := []int{2, 7, 11, 15}
	// target := 9

	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)

	fmt.Println("This is the answer: ", result)
}

func twoSum(nums []int, target int) []int {

	// step through each element in the array
	// subtract the target from the current value
	// look for the value in the hashmap
	// if it doesnt exists, add the curent element and index to the map
	// upon match return the 2 indicies

	stephensMap := make(map[int]int, len(nums))

	for currIndex, currValue := range nums {

		if requiredIndex, isPresent := stephensMap[target-currValue]; isPresent {
			return []int{requiredIndex, currIndex}
		}

		stephensMap[currValue] = currIndex
	}

	return []int{}
}
