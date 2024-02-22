package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// len == 0
	if len(preorder) == 0 {
		return nil
	}
	val := preorder[0]
	root := &TreeNode{Val: val}
	// len == 1
	if len(preorder) == 1 {
		return root
	}
	// len >= 2
	leftVal := preorder[1]
	i := 0
	for ; i < len(postorder)-1; i++ {
		if postorder[i] == leftVal {
			break
		}
	}
	root.Left = constructFromPrePost(preorder[1:1+i+1], postorder[:i+1])
	root.Right = constructFromPrePost(preorder[1+i+1:], postorder[i+1:len(postorder)-1])
	return root
}
