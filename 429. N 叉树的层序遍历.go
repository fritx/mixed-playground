package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	ans := make([][]int, 0)
	q := []*Node{root}
	for len(q) > 0 {
		ans = append(ans, []int{})
		//for _, n := range q {  // ** 1 bug
		// for _, n := range q[:] {
		//cq = make([]*Node, len(q))
		//copy(cq, q)  // ** 3 impr
		cq := q
		q = nil
		// q = []*Node{}
		for _, n := range cq {
			if n == nil { // ** 2 bug
				continue
			}
			q = append(q, n.Children...)
			ans[len(ans)-1] = append(ans[len(ans)-1], n.Val)
		}
	}
	if len(ans) > 0 && len(ans[len(ans)-1]) == 0 {
		ans = ans[:len(ans)-1]
	}
	return ans
}
