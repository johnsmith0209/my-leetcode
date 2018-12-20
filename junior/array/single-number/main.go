package main

import "fmt"

func singleNumber(nums []int) int {
	n := len(nums)
	rzt := 0
	for i := 0; i < n; i++ {
		rzt ^= nums[i]
	}
	return rzt
}

func main() {
	a := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(a))
}
