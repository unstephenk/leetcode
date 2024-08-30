package main

import "fmt"

func main() {
	// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
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
	if len(s) == 0 || len(s)%2 == 1 { // capture 2 edge cases
		return false
	}

	// create a mapping. Closed parenthesis maps to open
	charMap := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, c := range s {
		if _, isPresent := charMap[c]; isPresent { // add it to the stack
			stack = append(stack, c)
		} else if len(stack) == 0 || charMap[stack[len(stack)-1]] != c { // the stack if empty or char is not in the map, return false
			return false
		} else { // pop from the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0

}
