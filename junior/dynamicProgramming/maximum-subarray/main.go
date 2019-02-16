package main

import "fmt"

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2, -1, 0, 2, 4, -2, 10, -11, 10}))
}
