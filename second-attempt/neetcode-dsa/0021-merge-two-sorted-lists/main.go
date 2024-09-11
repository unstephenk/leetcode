package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Input: list1 = [1,2,4], list2 = [1,3,4]
	// Output: [1,1,2,3,4,4]

	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	res := mergeTwoLists(list1, list2)
	fmt.Println(res)

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

}
