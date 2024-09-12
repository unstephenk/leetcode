package main

func main() {
	// Design your implementation of the linked list. You can choose to use a singly or doubly linked list.
	// A node in a singly linked list should have two attributes: val and next. val is the value of the current node,
	// and next is a pointer/reference to the next node.
	// If you want to use the doubly linked list, you will need one more attribute prev to indicate the previous node
	// in the linked list. Assume all nodes in the linked list are 0-indexed.

	// ["MyLinkedList", "addAtHead", "addAtTail", "addAtIndex", "get", "deleteAtIndex", "get"]
	// [[], [1], [3], [1, 2], [1], [1], [1]]
}

type MyLinkedList struct {
}

func Constructor() MyLinkedList {

}

func (this *MyLinkedList) Get(index int) int {

}

func (this *MyLinkedList) AddAtHead(val int) {

}

func (this *MyLinkedList) AddAtTail(val int) {

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {

}

func (this *MyLinkedList) DeleteAtIndex(index int) {

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
