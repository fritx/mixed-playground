package trees

import (
	"container/heap"
	"playground/data_structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	queue := []*TreeNode{root}
	h := &data_structures.MaxHeap[int64]{}
	levels := 0
	for len(queue) > 0 {
		// ** bug 6.1
		// levels++
		// ** bug 2
		// node := queue[0]
		// queue = queue[1:]
		// if node == nil {
		// 	continue
		// }
		nextQueue := []*TreeNode{}
		var sum int64
		hasAny := false
		for _, n := range queue {
			if n == nil {
				continue
			}
			// ** bug 6.2
			hasAny = true
			// ** bug 3
			// queue = append(queue, n.Left)
			// queue = append(queue, n.Right)
			nextQueue = append(nextQueue, n.Left)
			nextQueue = append(nextQueue, n.Right)
			sum += int64(n.Val)
			// ** bug 4
			// heap.Push(h, sum)
		}
		// ** bug 6.2
		if hasAny {
			levels++
			heap.Push(h, sum)
		}
		queue = nextQueue
	}
	// ** bug 1
	// if k > h.Len() {
	// ** bug 5
	// if k > levels {
	// ** impr 1
	if k <= 0 || k > levels {
		return -1
	}
	var ans int64
	for i := 0; i < k; i++ {
		// ** bug 5
		// ans = h.Pop().(int64)
		ans = heap.Pop(h).(int64)
	}
	return ans
}
