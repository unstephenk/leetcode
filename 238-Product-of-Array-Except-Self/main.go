package main

import "fmt"

func main() {
	value := []int{1, 2, 3, 4}
	// value := [-1,1,0,-3,3]

	productExceptSelf(value)
}

func productExceptSelf(nums []int) []int {

	// scan the array and calculate our prefixes
	// add the prefix to the answer array
	// start at the end of the array and multiply our postfix to the answer

	answer := make([]int, len(nums))

	// start prefix at 1
	prefixPointer := 1

	// create prefix array
	for numIndex, numValue := range nums {
		answer[numIndex] = prefixPointer
		prefixPointer = prefixPointer * numValue
	}

	// start postfix at 1
	postfixPointer := 1

	// create postfix array
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] = answer[i] * postfixPointer
		postfixPointer = nums[i] * postfixPointer
	}

	fmt.Println(answer)

	return answer
}
