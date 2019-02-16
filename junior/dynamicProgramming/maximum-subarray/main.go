package main

import (
	"fmt"
	"math"
)

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

func maxSubArray2(nums []int) int {
	l, answer := len(nums), nums[0]
	dp := make([]int, l)
	dp[0] = nums[0]
	for i := 1; i < l; i++ {
		dp[i] = int(math.Max(float64(nums[i]+dp[i-1]), float64(nums[i])))
		answer = int(math.Max(float64(dp[i]), float64(answer)))
	}
	return answer
}

func main() {
	fmt.Println(maxSubArray([]int{-2, -1, 0, 2, 4, -2, 10, -11, 10}))
	fmt.Println(maxSubArray2([]int{-2, -1, 0, 2, 4, -2, 10, -11, 10}))
	// fmt.Println(maxSubArray3([]int{-2, -1, 0, 2, 4, -2, 10, -11, 10}))
}
