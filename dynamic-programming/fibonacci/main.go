/**
1. recursion
2. memoize
3. bottom up
*/
package main

import (
	"fmt"
	"time"
)

// 1. recursion
func fib1(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

// 2. memoize
func fib2(n int, memo map[int]int) int {
	if n == 1 || n == 2 {
		return 1
	}
	v, exists := memo[n]
	if exists {
		return v
	}
	v = fib2(n-1, memo) + fib2(n-2, memo)
	memo[n] = v
	return v
}

// 3. bottom up
func fib3(n int) int {
	memo := make(map[int]int)
	memo[1] = 1
	memo[2] = 1
	for i := 3; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func main() {
	start := time.Now()
	// O(2^n)
	fmt.Println(fib1(40))
	elapsed := time.Since(start)
	fmt.Printf("fib1 use time: %s", elapsed) // 526.690046ms
	// O(n)
	start = time.Now()
	memo := make(map[int]int)
	fmt.Println(fib2(40, memo))
	elapsed = time.Since(start)
	fmt.Printf("fib2 use time: %s", elapsed) // 23.867µs
	// O(n)
	start = time.Now()
	fmt.Println(fib3(40))
	elapsed = time.Since(start)
	fmt.Printf("fib3 use time: %s", elapsed) // 8.424µs
	// fib2和fib3时间上的区别就是recursion和iteration的区别了
}
