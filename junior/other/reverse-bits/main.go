package main

func missingNumber(nums []int) int {
	res, l := len(nums), len(nums)
	for i := 0; i < l; i++ {
		res ^= nums[i]
		res ^= i
	}
	return res
}

func main() {

}
