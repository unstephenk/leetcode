package main

import "fmt"

func main() {
	// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the
	// input string is valid.
	// An input string is valid if:
	// Open brackets must be closed by the same type of brackets.
	// Open brackets must be closed in the correct order.
	// Every close bracket has a corresponding open bracket of the same type.

	// s := "()"
	// s := "()[]{}"
	s := "(]"

	res := isValid(s)

	fmt.Println(res)
}

func isValid(s string) bool {
	// iterate through the string
	// check to see if the closing bracket is at the top of the stack
	// if so, remove from the stack and move on
	// if not, add the variable to the stack
	// at the end of the stack return the true if the stack is empty

	stack := []rune{} // create a stack

	sMap := map[rune]rune{ // create a map of the closures
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, sChar := range s {
		topOfStack := len(stack) - 1

		// check to make sure the char is an open bracket by checking to make sure it exists in our map
		if _, isPresent := sMap[sChar]; isPresent {
			stack = append(stack, sChar)
		} else if len(stack) == 0 || sMap[stack[topOfStack]] != sChar { // else this is a closing bracket and we need to find the opening bracket at the top of the stack
			return false
		} else {
			stack = stack[:topOfStack]
		}
	}

	return len(stack) == 0
}
