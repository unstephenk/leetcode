package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	// Given the head of a singly linked list, reverse the list, and return the reversed list.
	// head := []int{1, 2, 3, 4, 5}

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	res := reverseListIterative(head)
	fmt.Println(res)

}

func reverseListIterative(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		tempNext := curr.Next

		curr.Next = prev
		prev = curr
		curr = tempNext
	}

	return prev

}
