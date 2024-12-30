package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// Given the head of a singly linked list, reverse the list, and return the reversed list.

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	res := reverseListIterative(head)
	fmt.Println(res)

}

func reverseListIterative(head *ListNode) *ListNode {
	// start the prev node at null
	var prev *ListNode
	curr := head

	// while loop until nil
	for curr != nil {
		// set the curr.next to prev
		// set prev curr
		// set curr to curr.next
		tempCurrNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tempCurrNext
	}
	return prev
}
