package main

import "fmt"

func main() {
	// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
	// A node in a singly linked list should have two attributes: val and next. val is the value of the current node,
	// and next is a pointer/reference to the next node.
	// If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node
	// in the linked list. Assume all nodes in the linked list are 0-indexed.

	// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
	// [[], [1], [3], [1, 2], [1], [1], [1]]

	var myLinkedList MyLinkedList

	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	myLinkedList.Get(1)           // return 2
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	myLinkedList.Get(1)           // return 3

	fmt.Println(myLinkedList.head)
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	next *Node
	prev *Node
	Val  int
}

func Constructor() MyLinkedList {
	head := &Node{Val: -1, next: nil, prev: nil}
	tail := &Node{Val: -1, next: nil, prev: nil}

	head.next = tail
	tail.prev = head

	return MyLinkedList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}

	curr := this.head

	for i := 0; i <= index; i++ {
		curr = curr.next
	}

	return curr.Val

}

func (this *MyLinkedList) AddAtHead(val int) {

	newNode := &Node{
		next: this.head.next,
		prev: this.head,
		Val:  val,
	}

	this.head.next.prev = newNode
	this.head.next = newNode
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		next: this.tail,
		prev: this.tail.prev,
		Val:  val,
	}

	this.tail.prev.next = newNode
	this.tail.prev = newNode
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) { // add it before the indexed node
	if index > this.size {
		return
	}

	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	curr := this.head

	for i := 0; i < index; i++ {
		curr = curr.next
	}

	newNode := &Node{
		next: curr.next,
		prev: curr,
		Val:  val,
	}

	this.head.next = newNode
	this.head.next.prev = newNode
	this.size++

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}

	curr := this.head.next

	for index > 0 {
		curr = curr.next
		index--
	}
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	this.size--

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
