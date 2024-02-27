// 235. 二叉搜索树的最近公共祖先
// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 一次遍历
// 题目OJ不会出错 但增强的测试用例 不存在的节点会出错
// func bst_lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	node := root
// 	for node != nil {
// 		if node.Val > q.Val && node.Val > p.Val {
// 			node = node.Left
// 		} else if node.Val < q.Val && node.Val < p.Val {
// 			node = node.Right
// 		} else {
// 			// ** bug 1. missing break
// 			break
// 		}
// 	}
// 	return node
// }

// 两次遍历
func bst_lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	path_p, path_q := bst_findPath(root, p), bst_findPath(root, q)
	var node *TreeNode
	for i := 0; i < len(path_p) && i < len(path_q); i++ {
		if path_p[i] == path_q[i] {
			node = path_p[i]
		} else {
			break
		}
	}
	return node
}
func bst_findPath(root, p *TreeNode) []*TreeNode {
	path := []*TreeNode{}
	q := []*TreeNode{root}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}
		path = append(path, node)
		if node.Val < p.Val {
			q = append(q, node.Right)
		} else if node.Val > p.Val {
			q = append(q, node.Left)
		} else {
			// ** bug 2.1
			// break
			return path
		}
	}
	// ** bug 2.2
	// return path
	return nil
}
