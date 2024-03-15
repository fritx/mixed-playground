package trees

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

// 递归 + map存储父节点
func lowestCommonAncestor_map(root, p, q *TreeNode) *TreeNode {
	parent := map[int]*TreeNode{}
	visited := map[int]bool{}
	// ** impr 1.1
	// 题目OJ不会出错 但增强的测试用例 不存在的节点会出错
	// exists := map[int]bool{}
	if root == nil {
		return nil
	}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// exists[node.Val] = true
		if node.Left != nil {
			parent[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	for p != nil {
		// ** impr 1.2
		// 题目OJ不会出错 但增强的测试用例 不存在的节点会出错
		// if !exists[p.Val] {
		if p.Val != root.Val && parent[p.Val] == nil {
			return nil
		}
		visited[p.Val] = true
		p = parent[p.Val]
	}
	for q != nil {
		// ** impr 1.3
		// 题目OJ不会出错 但增强的测试用例 不存在的节点会出错
		// if !exists[q.Val] {
		if q.Val != root.Val && parent[q.Val] == nil {
			return nil
		}
		if visited[q.Val] {
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

// 层序遍历 + struct存储路径
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	x, y := findPath(root, p), findPath(root, q)
	if len(y) < len(x) {
		x, y = y, x // 令 x小 y大
	}
	// xl, yl := len(x), len(y)
	// ** impr 3
	// 题目OJ不会出错 但增强的测试用例 边界条件空x,y会出错
	// if xl == 0 {
	// 	return nil
	// }
	// for i := 0; i < yl-xl; i++ {
	// 	yn := y[yl-1-i]
	// 	if yn == x[xl-1] {
	// 		return yn
	// 	}
	// }
	// for i := 0; i < xl; i++ {
	// 	a, b := x[xl-1-i], y[xl-1-i]
	// 	if a == b {
	// 		return a
	// 	}
	// }
	// return nil
	// ** impr 4. simpler
	var ans *TreeNode
	for i := 0; i < len(x); i++ {
		if x[i] == y[i] {
			ans = x[i]
		} else {
			break
		}
	}
	return ans
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
			// ** bug 3.opt1 - 超出时间限制
			// cPath := make([]*TreeNode, len(item.Path))
			// copy(cPath, item.Path)
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(cPath, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(cPath, item.Node.Right)})
			// ** bug 3.opt2
			// x, y := append([]*TreeNode{}, item.Path...), append([]*TreeNode{}, item.Path...)
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(x, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(y, item.Node.Right)})
			// ** bug 3.opt3 - requires Go >=1.22
			// x, y := slices.Concat[[]*TreeNode]([]*TreeNode{}, item.Path), slices.Concat[[]*TreeNode]([]*TreeNode{}, item.Path)
			// nxq = append(nxq, &QueueItem{item.Node.Left, append(x, item.Node.Left)})
			// nxq = append(nxq, &QueueItem{item.Node.Right, append(y, item.Node.Right)})
			// ** bug 3.opt4
			path := item.Path[:len(item.Path):len(item.Path)]
			nxq = append(nxq, &QueueItem{item.Node.Left, append(path, item.Node.Left)})
			nxq = append(nxq, &QueueItem{item.Node.Right, append(path, item.Node.Right)})
		}
		q = nxq
	}
	return []*TreeNode{}
}
