package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归 / O(n),O(n)
func preorderTraversal_dfs(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// 迭代 / O(n),O(n)
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

// 迭代 2 / O(n),O(n)
func preorderTraversal_2(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	curr := root

	for len(stack) > 0 || curr != nil {
		if curr != nil {
			ans = append(ans, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		} else {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			curr = curr.Right
		}
	}
	return ans
}

// todo: Morris 遍历 / O(n),O(1)
