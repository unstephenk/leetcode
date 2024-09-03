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

	charMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, c := range s {
		if _, isPresent := charMap[c]; isPresent { // add it to the stack
			stack = append(stack, c)
		} else if len(stack) == 0 || charMap[stack[len(stack)-1]] != c { // if the stack is empty or if c does not match the end of the stack
			return false
		} else { // pop from the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
