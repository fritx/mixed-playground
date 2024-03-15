package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, n.Val)
		curr = n.Right
	}
	return ans
}
