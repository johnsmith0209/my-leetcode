package main

import "fmt"

func plusOne(digits []int) []int {
	n, stepOn := len(digits), true
	for i := n - 1; i >= 0; i-- {
		if stepOn {
			stepOn = false
			digits[i]++
		}
		if digits[i] == 10 {
			digits[i] = 0
			stepOn = true
		} else {
			break
		}
	}
	if stepOn {
		fmt.Println(digits)
		digits[0] = 0
		digits = append([]int{1}, digits...)
		fmt.Println(digits)
	}
	return digits
}

func main() {
	a := []int{9, 9, 9}
	// fmt.Println(plusOne(a))
	fmt.Println(plusOne(plusOne(a)))
}
