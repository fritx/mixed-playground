package main

import (
	"encoding/json"
	"fmt"
	"playground/utils"
	"testing"
)

func Test_bst_lowestCommonAncestor(t *testing.T) {
	cases := []any{}
	// 题目自身约束
	// 2 <= n <= 10^5
	// expected 'root' to have 2 <= size <= 100000 but got 1
	if err := json.Unmarshal([]byte(`[
		[6,2,8,0,4,7,9,null,null,3,5], 2, 8, 6,
		[6,2,8,0,4,7,9,null,null,3,5], 2, 4, 2,
		[6,2,8,0,4,7,9,null,null,3,5], 2, 4, 2,
		[2,1], 2, 1, 2,
		[1], 1, 1, 1,
		[1], 1, 2, null,
		[], 1, 2, null,
		[], 1, 1, null
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	us := 4
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		data := utils.SliceToIntIfFloat64(cases[us*i].([]any))
		var p, q, want *TreeNode
		if v, ok := cases[us*i+1].(float64); ok {
			p = &TreeNode{Val: int(v)}
		}
		if v, ok := cases[us*i+2].(float64); ok {
			q = &TreeNode{Val: int(v)}
		}
		if v, ok := cases[us*i+3].(float64); ok {
			want = &TreeNode{Val: int(v)}
		}
		testEach_bst_lowestCommonAncestor(t, bst_lowestCommonAncestor, data, p, q, want)
	}
}

func testEach_bst_lowestCommonAncestor(t *testing.T, fn func(root, p, q *TreeNode) *TreeNode, data []any, p, q, want *TreeNode) {
	root := sliceToBinaryTree(data)
	ans := fn(root, p, q)

	if ans == nil || want == nil {
		utils.AssertEqual(t, ans, want,
			fmt.Sprintf("Not equal. data=%v, p=%v, q=%v", data, p, q))
	} else {
		utils.AssertEqual(t, ans.Val, want.Val,
			fmt.Sprintf("Not equal. data=%v, p=%v, q=%v", data, p, q))
	}
}
