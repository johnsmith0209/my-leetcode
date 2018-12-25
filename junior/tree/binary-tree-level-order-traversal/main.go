package main

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
func levelOrder(root *TreeNode) [][]int {
	s, children := [][]int{}, []*TreeNode{}
	if root == nil {
		return s
	}
	for len(children) > 0 {
		nums, newChildren := []int{}, []*TreeNode{}
		for _, v := range children {
			nums = append(nums, v.Val)
			if v.Left != nil {
				newChildren = append(newChildren, v.Left)
			}
			if v.Right != nil {
				newChildren = append(newChildren, v.Right)
			}
		}
		s, children = append(s, nums), newChildren
	}
	return s
}

func main() {

}
