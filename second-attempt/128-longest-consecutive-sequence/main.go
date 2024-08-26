package main

import "fmt"

func main() {

	nums := []int{100, 4, 200, 1, 3, 2}

	res := longestConsecutive(nums)

	fmt.Println(res)

}

func longestConsecutive(nums []int) int {

	stephensMap := make(map[int]bool, len(nums))

	for _, numsValue := range nums {
		stephensMap[numsValue] = true
	}

	res := 0

	for _, numsValue := range nums {
		if stephensMap[numsValue+1] { // skip this iteration since its not the start of a sequence
			continue
		}

		cnt := 0 // count the number of digits in a sequence

		for i := numsValue; stephensMap[i]; i-- { // if a digit to the left exists, its a sequence
			cnt++
		}

		if cnt > res {
			res = cnt
		}
	}

	return res
}