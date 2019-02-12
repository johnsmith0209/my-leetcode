package main

import "fmt"

func hammingWeight(num uint32) int {
	rzt := 0
	for ; num > 0; num = num >> 1 {
		if num&1 == 1 {
			rzt++
		}
	}
	return rzt
}

func main() {
	fmt.Println(hammingWeight(123))
}
