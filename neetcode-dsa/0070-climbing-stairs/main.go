package main

import "fmt"

func main() {
	// You are climbing a staircase. It takes n steps to reach the top.
	// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

	inputNum := 2

	res := climbStairs(inputNum)
	fmt.Println(res)
}

func climbStairs(n int) int {
	one, two := 1, 1

	for i := 0; i < n-1; i++ {
		temp := one
		one = one + two
		two = temp
	}

	return one
}
