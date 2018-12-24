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
func flatTree(root *TreeNode, rzt []int) []int {
	if root == nil {
		return rzt
	}
	if root.Left != nil {
		rzt = flatTree(root.Left, rzt)
	}
	rzt = append(rzt, root.Val)
	if root.Right != nil {
		rzt = flatTree(root.Right, rzt)
	}
	return rzt
}

func isValidBST(root *TreeNode) bool {
	arr := flatTree(root, []int{})
	for i := 0; i < len(arr)-1; i++ {
		cur, next := arr[i], arr[i+1]
		if cur >= next {
			return false
		}
	}
	return true
}

func main() {
	root := TreeNode{Val: 1}
	a1 := TreeNode{Val: 2}
	a2 := TreeNode{Val: 3}
	a3 := TreeNode{Val: 4}
	root.Left = &a1
	a1.Left = &a2
	a1.Right = &a3
	fmt.Println(flatTree(&root, []int{}))
}
