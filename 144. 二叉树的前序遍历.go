package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == nil {
			continue
		}
		stack = append(stack, curr.Right)
		stack = append(stack, curr.Left)
		ans = append(ans, curr.Val)
	}
	return ans
}
