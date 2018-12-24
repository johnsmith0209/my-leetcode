package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := maxDepth(root.Left)+1, maxDepth(root.Right)+1
	if left > right {
		return left
	} else {
		return right
	}
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a[0], a[1:])

	root := TreeNode{Val: 1}
	a1 := TreeNode{Val: 2}
	// a2 := TreeNode{Val: 3}
	// a3 := TreeNode{Val: 4}
	root.Left = &a1
	// a1.Left = &a2
	// a1.Right = &a3
	fmt.Println(maxDepth(&root))
}
