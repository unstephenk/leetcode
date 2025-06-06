package main

import "fmt"

func main() {

	value := []int{1, 1, 1, 2, 2, 3}
	k := 2

	// value := [1]int{1}
	// k := 1
	topKFrequent(value, k)
}

func topKFrequent(nums []int, k int) []int {
	// return the top most frequent elements
	// use bucket sort

	// create a map with the element in the array and update its values

	// create the bucket as the length of the array
	// add the values from the map

	// create an array for the answer with length 0 and capacity of k

	stephensMap := make(map[int]int, len(nums))

	// elementValue -> frequency
	for _, value := range nums {
		stephensMap[value]++
	}

	stephensBucket := make([][]int, len(nums)+1)

	// [[], [3], [2], [1]]
	for key, value := range stephensMap {
		stephensBucket[value] = append(stephensBucket[value], key)
	}

	// convert the bucket to the answer => [1, 2]
	stephensAnswer := make([]int, 0, len(nums))

	for i := len(stephensBucket) - 1; i >= 0; i-- {
		for _, value := range stephensBucket[i] {
			// use k
			if k > 0 {
				stephensAnswer = append(stephensAnswer, value)
				k--
			}
		}
	}

	fmt.Println(stephensAnswer)
	return stephensAnswer

}
