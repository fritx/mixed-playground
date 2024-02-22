package main

// 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createBinaryTreeFromSlice creates a binary tree from a slice of integers
// Assuming the slice is completely filled with integers, starting with the root at index 0
// func createBinaryTreeFromSlice(slice []interface{}) *TreeNode {
// 	if len(slice) == 0 {
// 		return nil
// 	}
// 	if slice[0] == nil {
// 		return nil
// 	}
// 	root := &TreeNode{Val: slice[0].(int)}

//		// If there's only one element, return the root
//		if len(slice) == 1 {
//			return root
//		}
//		// Recursively create the left and right subtrees
//		root.Left = createBinaryTreeFromSlice(slice[1 : (len(slice)+1)/2])
//		root.Right = createBinaryTreeFromSlice(slice[(len(slice)+1)/2:])
//		return root
//	}

// 需要适配LeetCode的二叉树数组表示形式！
// @author Bingo
func sliceToBinaryTree(nums []interface{}) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	root := &TreeNode{Val: nums[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if nums[i] != nil {
			node.Left = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != nil {
			node.Right = &TreeNode{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// @author 文心一言
// traverseBinaryTree traverses the binary tree and returns a slice of int or nil values
func traverseBinaryTree(root *TreeNode) []interface{} {
	if root == nil {
		return nil
	}

	// Create a slice to store the values and nils
	result := make([]interface{}, 0)

	// Preorder traversal: root -> left -> right
	result = append(result, root.Val)
	result = append(result, traverseBinaryTree(root.Left)...)
	result = append(result, traverseBinaryTree(root.Right)...)

	// Check if the left child is nil and the right child is not nil
	// If so, insert a nil value to represent the null left child
	if root.Left == nil && root.Right != nil {
		result = append(result, nil)
	}

	return result
}

// @author Copilot
func createTree(arr []interface{}) *Node {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := &Node{Val: arr[0].(int)}
	queue := []*Node{root}
	i := 2

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for i < len(arr) && arr[i] != nil {
			child := &Node{Val: arr[i].(int)}
			node.Children = append(node.Children, child)
			queue = append(queue, child)
			i++
		}

		i++
	}

	return root
}

// ========================
// N叉树
type Node struct {
	Val      int
	Children []*Node
}

// @author Bingo
func sliceToNTree(nums []interface{}) *Node {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}

	root := &Node{Val: nums[0].(int)}
	queue := []*Node{root}
	i := 1

	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		for i < len(nums) && nums[i] != nil {
			child := &Node{Val: nums[i].(int)}
			node.Children = append(node.Children, child)
			queue = append(queue, child)
			i++
		}
		if i < len(nums) && nums[i] == nil {
			i++ // Skip the nil marking the end of a line
		}
	}

	return root
}

// @author 文心一言
// convertSliceToNaryTree converts a slice of integers or nils to an N-ary tree
func convertSliceToNaryTree(slice []interface{}) *Node {
	if len(slice) == 0 {
		return nil
	}

	// Create the root node
	root := &Node{Val: slice[0].(int)}

	// Iterate over the rest of the slice to build the tree
	children := make([]*Node, 0)
	for i := 1; i < len(slice); i++ {
		// Check if the current element is nil, indicating a new level
		if slice[i] == nil {
			// If children are present, add them to the previous node
			if len(children) > 0 {
				prevNode := children[len(children)-1]
				prevNode.Children = children
				children = make([]*Node, 0) // Reset children slice for the next level
			}
		} else {
			// Create a new child node
			newNode := &Node{Val: slice[i].(int)}
			children = append(children, newNode)
		}
	}

	// If there are remaining children, set them as the last node's children
	if len(children) > 0 {
		lastNode := children[len(children)-1]
		lastNode.Children = children
	}

	return root
}
