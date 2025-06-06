package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// You are given the heads of two sorted linked lists list1 and list2.
	// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
	// Return the head of the merged linked list.

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}

	res := mergeTwoLists(list1, list2)
	fmt.Println(res)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2

}
