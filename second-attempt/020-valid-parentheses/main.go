package main

import "fmt"

func main() {
	// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
	// An input string is valid if:
	// Open brackets must be closed by the same type of brackets.
	// Open brackets must be closed in the correct order.
	// Every close bracket has a corresponding open bracket of the same type.

	// s := "()"
	s := "()[]{}"
	// s := "(]"

	res := isValid(s)
	fmt.Println(res)

}

func isValid(s string) bool {
	// edge case
	// create the character map
	// create a stack
	// step through the string

	if len(s) == 0 || len(s)%2 == 1 { // capture 2 edge cases
		return false
	}

	stephensMap := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := []rune{}

	// peek at i, check the character map
	// if in char map, grab the value and check the top of the stack
	// if its in the top of the stack, recreate the stack by removing the top

	for _, sChar := range s {
		if _, isPresent := stephensMap[sChar]; isPresent {
			stack = append(stack, sChar)
		} else if len(stack) == 0 || stephensMap[stack[len(stack)-1]] != sChar {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
