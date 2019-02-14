package main

import "fmt"

func climbStairs(n int) int {
	result, preResult, prePreResult := 1, 1, 0
	for n >= 2 {
		n--
		preResult, prePreResult = result, preResult
		result = prePreResult + preResult
	}
	return result
}

func main() {
	fmt.Println(44, climbStairs(44))
	// fmt.Println(1, climbStairs(1))
	// fmt.Println(2, climbStairs(2))
	// fmt.Println(3, climbStairs(3))
	// fmt.Println(4, climbStairs(4))
	// fmt.Println(5, climbStairs(5))
	// fmt.Println(6, climbStairs(6))
	// fmt.Println(7, climbStairs(7))
}
