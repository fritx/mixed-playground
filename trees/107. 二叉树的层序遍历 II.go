package trees

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	q := []*TreeNode{root}

	for len(q) > 0 {
		ans = append(ans, []int{})
		// for i := 0; i < len(q); i++ { // **
		n := len(q)
		for i := 0; i < n; i++ {
			t := q[0]
			q = q[1:]
			if t == nil {
				continue
			}
			q = append(q, t.Left)
			q = append(q, t.Right)
			ans[len(ans)-1] = append(ans[len(ans)-1], t.Val)
		}
	}
	if len(ans) > 0 && len(ans[len(ans)-1]) == 0 {
		// ans = ans[1:] // **
		ans = ans[:len(ans)-1]
	}
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}
