package main

import "fmt"

func main() {

	// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	result := topKFrequent(nums, k)

	fmt.Println(result)
}

func topKFrequent(nums []int, k int) []int {

	// create a map and use the nums value as the key
	// increment the value of the key everytime we find a dup
	// convert the first map to a Bucket Sort
	// start at the last element in the bucket sort and find the k most frequent values

	frequencyMap := make(map[int]int, len(nums))

	// { nums: frequency }
	for _, numsValue := range nums {
		frequencyMap[numsValue]++
	}

	// [[0], [numsValue], [numsValue]]
	bucketArray := make([][]int, len(nums)+1)

	for numsValue, frequencyCount := range frequencyMap {
		bucketArray[frequencyCount] = append(bucketArray[frequencyCount], numsValue)
	}

	answerArray := make([]int, 0, len(nums))

	for i := len(bucketArray) - 1; i >= 0; i-- {
		for _, bucketValue := range bucketArray[i] {
			if k > 0 {
				answerArray = append(answerArray, bucketValue)
				k--
			}
		}
	}

	return answerArray

}
