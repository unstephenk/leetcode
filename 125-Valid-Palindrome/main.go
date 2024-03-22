package main

import "fmt"

func main() {
	value := "A man, a plan, a canal: Panama"

	res := isPalindrome(value)

	fmt.Println(res)
}

func isPalindrome(s string) bool {
	// create a function to check punctuation
	// create a function to convert to lowercase
	// step through the array and compare the first element to the last element
	// if match continue
	// if pointers end up at same location, stop and return answer

	leftPointer := 0
	rightPointer := len(s) - 1

	for leftPointer < rightPointer {
		if isPunctuation(s[leftPointer]) {
			leftPointer++
			continue
		}

		if isPunctuation(s[rightPointer]) {
			rightPointer--
			continue
		}

		if toLowerCase(s[leftPointer]) == toLowerCase(s[rightPointer]) {
			leftPointer++
			rightPointer--
		} else {
			return false
		}
	}

	return true

}

func toLowerCase(c byte) byte {
	if c >= 65 && c <= 90 {
		return c + 32
	}
	return c
}

func isPunctuation(c byte) bool {
	if (c >= 32 && c <= 47) || (c >= 58 && c <= 64) || (c >= 91 && c <= 96) || (c >= 123 && c <= 126) {
		return true
	}
	return false
}
