package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Given the head of a singly linked list, reverse the list, and return the reversed list.

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	res := reverseList(head)
	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {

	// set prev and curr
	var prev *ListNode
	curr := head

	for curr != nil {
		tempCurrNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tempCurrNext
	}

	return prev

}
