package main

import (
	"fmt"
	"math"
)

func rob(nums []int) int {
	l := len(nums)
	dp := make([]int, l)
	dp[0] = nums[0]
	dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	for i := 2; i < l; i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+nums[i]), float64(dp[i-1])))
	}
	fmt.Println(dp)
	return dp[l-1]
}

func main() {
	fmt.Println(rob([]int{2, 1, 1, 4}))
	fmt.Println(rob([]int{2, 1, 1, 2, 4}))
}
