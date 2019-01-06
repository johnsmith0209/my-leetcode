package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	if l == 1 {
		return &TreeNode{Val: nums[0]}
	}
	mid := l / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  buildTree(nums[:mid]),
		Right: buildTree(nums[mid+1:]),
	}
	fmt.Println("mid is", mid, nums[:mid], nums[mid+1:])
	return root
}

func buildTree2(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}
	mid := (l + r) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  buildTree(nums, l, mid-1),
		Right: buildTree(nums, mid+1, r),
	}
	return root
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return buildTree(nums)
}

func main() {

}
