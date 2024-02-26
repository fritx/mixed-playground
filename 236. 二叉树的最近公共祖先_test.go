package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	cases := []any{}
	// 题目自身约束
	// 2 <= n <= 10^5
	// expected 'root' to have 2 <= size <= 100000 but got 1
	if err := json.Unmarshal([]byte(`[
		[3,5,1,6,2,0,8,null,null,7,4], 5, 1, 3,
		[3,5,1,6,2,0,8,null,null,7,4], 5, 4, 5,
		[3,5,1,6,2,0,8,null,null,7,4], 3, 5, 3,
		[1,2], 1, 2, 1,
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
		data := sliceToIntIfFloat64(cases[us*i].([]any))
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
		// testEach_lowestCommonAncestor(t, lowestCommonAncestor, data, p, q, want)
		testEach_lowestCommonAncestor(t, lowestCommonAncestor_dfs, data, p, q, want)
	}
}

func testEach_lowestCommonAncestor(t *testing.T, fn func(root, p, q *TreeNode) *TreeNode, data []any, p, q, want *TreeNode) {
	root := sliceToBinaryTree(data)
	ans := fn(root, p, q)

	if ans == nil || want == nil {
		AssertEqual[*TreeNode](t, ans, want,
			fmt.Sprintf("Not equal. data=%v, p=%v, q=%v", data, p, q))
	} else {
		AssertEqual[int](t, ans.Val, want.Val,
			fmt.Sprintf("Not equal. data=%v, p=%v, q=%v", data, p, q))
	}
}
