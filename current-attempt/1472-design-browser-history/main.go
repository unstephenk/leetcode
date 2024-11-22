package main

import "fmt"

func main() {
	// You have a browser of one tab where you start on the homepage and you can visit another url, get back in the history number
	// of steps or move forward in the history number of steps.
	//
	// Implement the BrowserHistory class:
	//
	// BrowserHistory(string homepage) Initializes the object with the homepage of the browser.
	// void visit(string url) Visits url from the current page. It clears up all the forward history.
	// string back(int steps) Move steps back in history. If you can only return x steps in the history and steps > x, you will return only x steps.
	// Return the current url after moving back in history at most steps.
	// string forward(int steps) Move steps forward in history. If you can only forward x steps in the history and steps > x, you will forward
	// only x steps. Return the current url after forwarding in history at most steps.

	// ["BrowserHistory","visit","visit","visit","back","back","forward","visit","forward","back","back"]
	history := Constructor("leetcode.com")
	fmt.Println(history)
	history.Visit("google.com")
	fmt.Println(history)
	history.Visit("facebook.com")
	fmt.Println(history)
	history.Visit("youtube.com")
	fmt.Println(history)
	history.Back(1)
	fmt.Println(history)
	history.Back(1)
	fmt.Println(history)
	history.Forward(1)
	fmt.Println(history)
	history.Visit("linkedin.com")
	fmt.Println(history)
	history.Forward(2)
	fmt.Println(history)
	history.Back(2)
	fmt.Println(history)
	history.Back(7)
	fmt.Println(history)
}

type BrowserHistory struct {
	head *linkedNode
}

type linkedNode struct {
	url  string
	prev *linkedNode
	next *linkedNode
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		head: &linkedNode{
			url: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	node := &linkedNode{
		url:  url,
		prev: this.head,
	}

	this.head.next = node
	this.head = node
}

func (this *BrowserHistory) Back(steps int) string {
	cur := this.head
	var prev *linkedNode
	for cur != nil && steps > 0 {
		prev = cur
		cur = cur.prev
		steps--
	}

	if cur != nil && steps == 0 {
		this.head = cur
		return cur.url
	}

	this.head = prev
	return prev.url
}

func (this *BrowserHistory) Forward(steps int) string {
	cur := this.head
	var next *linkedNode
	for cur != nil && steps > 0 {
		next = cur
		cur = cur.next
		steps--
	}

	if cur != nil && steps == 0 {
		this.head = cur
		return cur.url
	}

	this.head = next
	return next.url
}
