package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代（不使用倒序）
// 图解！玩转二叉树遍历（非递归）
// https://zhuanlan.zhihu.com/p/577745119
func postorderTraversal_2(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			// 当前节点入栈
			stack = append(stack, curr)
			// 如果当前节点有左子树，继续向左子树找
			if curr.Left != nil {
				curr = curr.Left
			} else {
				// 如果当前节点无左子树，在右子树继续找
				curr = curr.Right
			}
		}
		// 跳出循环的条件是 root 为空，那当前栈顶元素为叶子节点。
		// 弹出栈顶元素，并加入结果数组
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		// 如果栈不为空，且当前栈顶元素的左节点是刚刚跳出的栈顶元素 cur
		// 则转向遍历右子树当前栈顶元素的右子树
		if len(stack) > 0 && stack[len(stack)-1].Left == curr {
			curr = stack[len(stack)-1].Right
		} else {
			curr = nil
		}
	}
	return ans
}

// 迭代（使用倒序）
func postorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curr == nil {
			continue
		}
		stack = append(stack, curr.Left)
		stack = append(stack, curr.Right)
		ans = append(ans, curr.Val)
	}
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}
