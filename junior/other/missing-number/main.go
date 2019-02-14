package main

import "fmt"

func missingNumber(nums []int) int {
	l, ext := len(nums), make([]int, len(nums)+1)
	for i := 0; i < l; i++ {
		if i == 0 || ext[i] == 0 {
			ext[i] = -1
		}
		ext[nums[i]] = nums[i]
	}
	for i := 0; i < len(ext); i++ {
		if i != ext[i] {
			return i
		}
	}
	return len(nums)
}

func missingNumber2(nums []int) int {
	res, l := len(nums), len(nums)
	for i := 0; i < l; i++ {
		fmt.Println(res, nums[i], res^nums[i])
		res ^= nums[i]
		fmt.Println(res, i, res^i)
		res ^= i
	}
	return res
}

func main() {
	// a := make([]int, 10)
	// a[2] = 10
	fmt.Println(missingNumber2([]int{0, 1, 2, 4}))
}
