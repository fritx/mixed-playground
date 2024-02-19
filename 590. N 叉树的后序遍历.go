package main

import "fmt"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 迭代（不使用倒序）
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/n-ary-tree-postorder-traversal/solutions/1327327/n-cha-shu-de-hou-xu-bian-li-by-leetcode-txesi/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func postorder_it(root *Node) (ans []int) {
	if root == nil {
		return
	}
	st := []*Node{}
	nextIndex := map[*Node]int{}
	node := root
	for len(st) > 0 || node != nil {
		for node != nil {
			st = append(st, node)
			if len(node.Children) == 0 {
				break
			}
			nextIndex[node] = 1
			node = node.Children[0]
		}
		node = st[len(st)-1]
		i := nextIndex[node]
		if i < len(node.Children) {
			nextIndex[node] = i + 1
			node = node.Children[i]
		} else {
			ans = append(ans, node.Val)
			st = st[:len(st)-1]
			delete(nextIndex, node)
			node = nil
		}
	}
	return
}

// 递归
func postorder(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, ch := range node.Children {
			if ch != nil {
				fmt.Println("ch.Val", ch.Val)
			}
			dfs(ch)
		}
		fmt.Println("node.Val", node.Val)
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}
