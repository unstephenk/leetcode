package main

import "fmt"

func main() {
	// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Println(res)
}

func topKFrequent(nums []int, k int) []int {
	// k = 2, return the top two most frequent elements

	// traverse the array and add the element as the key and the freq as the value
	// once we have our freq map, create a bucket sort array and add the elements to the bucket
	// finally start at the end of the bucket and return the k most frequent elements

	freqMap := map[int]int{}
	for _, num := range nums {
		freqMap[num]++
	}

	// {1: 4, 2: 10}

	bucketSortArray := []int{}
	for freqValue, freqCount := range freqMap {
		bucketSortArray[freqCount] = append(bucketSortArray[freqCount], freqValue)
	}

}
