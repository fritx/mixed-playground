// 938. 二叉搜索树的范围和
// https://leetcode.cn/problems/range-sum-of-bst/
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
func rangeSumBST_it(root *TreeNode, low int, high int) int {
	sum := 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		if node.Val < low {
			q = append(q, node.Right)
		} else if node.Val > high {
			q = append(q, node.Left)
		} else {
			q = append(q, node.Left)
			q = append(q, node.Right)
			sum += node.Val
		}
	}
	return sum
}

// 递归
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	// sum := 0
	// if root.Val >= low && root.Val <= high {
	//     sum += root.Val
	// }
	// sum += rangeSumBST(root.Left, low, high)
	// sum += rangeSumBST(root.Right, low, high)
	// ** patch 1.1
	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}
	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}
	// return sum
	// ** patch 1.2
	return root.Val + rangeSumBST(root.Left, low, high) +
		rangeSumBST(root.Right, low, high)
}
