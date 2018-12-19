package main

import "fmt"

func rotate(nums []int, k int) {
	if len(nums) <= 1 {
		return
	}
	fmt.Println(k, len(nums), k%len(nums))
	k = k % len(nums)
	for i := 0; i < k; i++ {
		cur := nums[0]
		for j := 1; j < len(nums); j++ {
			temp := nums[j]
			fmt.Println(temp, nums[j], cur)
			nums[j] = cur
			cur = temp
		}
		nums[0] = cur
	}
}

func rotate2(nums []int, k int) {
	n := len(nums)
	if len(nums) < 2 {
		return
	}
	k = k % len(nums)
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	arr := []int{1, 2}
	rotate2(arr, 1)
	// reverse(arr, 0, 3)
	fmt.Println(arr)
}
