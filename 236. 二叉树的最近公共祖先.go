package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
// func lowestCommonAncestor_dfs(root, p, q *TreeNode) *TreeNode {
// 	// ** impr 1
// 	// 题目OJ不会出错 但我们测试构造节点的方式下会出错
// 	// if root == nil || root == p || root == q {
// 	if root == nil || p == nil || q == nil ||
// 		root.Val == p.Val || root.Val == q.Val {
// 		return root
// 	}
// 	left := lowestCommonAncestor_dfs(root.Left, p, q)
// 	right := lowestCommonAncestor_dfs(root.Right, p, q)
// 	if left == nil {
// 		return right
// 	}
// 	if right == nil {
// 		return left
// 	}
// 	return root
// }

// 递归
// ** impr 2
// 题目OJ不会出错 但增强的测试用例 不存在的节点会出错
type ResultItem struct {
	Node       *TreeNode
	hasP, hasQ bool
}

func lowestCommonAncestor_dfs(root, p, q *TreeNode) *TreeNode {
	var dfs func(root, p, q *TreeNode) *ResultItem
	dfs = func(root, p, q *TreeNode) *ResultItem {
		res := &ResultItem{Node: root}
		if root == nil || p == nil || q == nil {
			return res
		}
		left := dfs(root.Left, p, q)
		if left.hasP && left.hasQ {
			return left
		}
		right := dfs(root.Right, p, q)
		if right.hasP && right.hasQ {
			return right
		}
		if root.Val == p.Val || left.hasP || right.hasP {
			res.hasP = true
		}
		if root.Val == q.Val || left.hasQ || right.hasQ {
			res.hasQ = true
		}
		return res
	}
	res := dfs(root, p, q)
	if res.hasP && res.hasQ {
		return res.Node
	}
	return nil
}

// 层序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x, y := findPath(root, p), findPath(root, q)
	if len(y) < len(x) {
		x, y = y, x // 令 x小 y大
	}
	xl, yl := len(x), len(y)
	// ** impr 3
	// 题目OJ不会出错 但增强的测试用例 边界条件空x,y会出错
	if xl == 0 {
		return nil
	}
	for i := 0; i < yl-xl; i++ {
		yn := y[yl-1-i]
		if yn == x[xl-1] {
			return yn
		}
	}
	for i := 0; i < xl; i++ {
		a, b := x[xl-1-i], y[xl-1-i]
		if a == b {
			return a
		}
	}
	return nil
}

type QueueItem struct {
	Node *TreeNode
	Path []*TreeNode
}

func findPath(root, p *TreeNode) []*TreeNode {
	// ** impr 2.1
	// if root == nil {
	if root == nil || p == nil {
		return []*TreeNode{}
	}
	// ** impr 1
	// redundant type from array, slice, or map composite literalsimplifycompositelit(default)
	// q := []*QueueItem{&QueueItem{Node: root, Path: []*TreeNode{root}}}
	q := []*QueueItem{{Node: root, Path: []*TreeNode{root}}}
	for len(q) > 0 {
		nxq := []*QueueItem{}
		for _, item := range q {
			if item.Node == nil {
				continue
			}
			// ** impr 2.2
			// if item.Node == p {
			if item.Node.Val == p.Val {
				return item.Path
			}
			// q = append(q, &QueueItem{item.Node.Left, append(item.Path, item.Node.Left)})
			// q = append(q, &QueueItem{item.Node.Right, append(item.Path, item.Node.Right)})
			// ** bug 2
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(item.Path, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(item.Path, item.Node.Right)})
			// ** bug 3.opt1
			// cPath := make([]*TreeNode, len(item.Path))
			// copy(cPath, item.Path)
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(cPath, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(cPath, item.Node.Right)})
			// ** bug 3.opt2
			// x, y := append([]*TreeNode{}, item.Path...), append([]*TreeNode{}, item.Path...)
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(x, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(y, item.Node.Right)})
			// ** bug 3.opt3
			path := item.Path[:len(item.Path):len(item.Path)]
			nxq = append(nxq, &QueueItem{item.Node.Left, append(path, item.Node.Left)})
			nxq = append(nxq, &QueueItem{item.Node.Right, append(path, item.Node.Right)})
		}
		q = nxq
	}
	return []*TreeNode{}
}
