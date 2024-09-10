package main

import "fmt"

func main() {

	// testCase := []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"}

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	minStack.GetMin() // return -3
	minStack.Pop()
	minStack.Top()    // return 0
	minStack.GetMin() // return -2

	fmt.Println(minStack)

}

type MinStack struct {
	stack    []int
	minStack []int
	top      int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}, -1}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 { // if this is the first time
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, min(val, (this.minStack)[this.top]))
	}
	this.stack = append(this.stack, val)
	this.top++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.minStack = this.minStack[:this.top]
	this.top--
	return
}

func (this *MinStack) Top() int {
	return (this.stack)[this.top]
}

func (this *MinStack) GetMin() int {
	return (this.minStack)[this.top]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
