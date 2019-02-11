package main

import "fmt"

func countPrimes(n int) int {
	record, stop, cnt := make([]bool, n+1), n/2, 0
	for i := 2; i <= stop; i++ {
		if !record[i] {
			for j := i + i; j < n; j += i {
				record[j] = true
			}
		}
	}
	for i := 2; i < n; i++ {
		if !record[i] {
			cnt++
		}
	}
	return cnt
}

func main() {
	fmt.Println("count prime 10", countPrimes(10))
}
