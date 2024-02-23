package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_kthLargestLevelSum(t *testing.T) {
	cases := []any{}
	// 题目自身约束
	// 2 <= n <= 10^5
	// expected 'root' to have 2 <= size <= 100000 but got 0
	// 1 <= k <= n
	// expected 'k' to have value from 1 to xxx only
	json.Unmarshal([]byte(`[
		[5,8,9,2,1,3,7,4,6], 2, 13,
		[1,2,null,3], 1, 3,
		[], 0, -1,
		[1], 0, -1,
		[1, 2], 0, -1,
		[], 1, -1,
		[1], 1, 1,
		[1, 2], 1, 2,
		[5,8,9,2,1,3,7], 4, -1
	]`), &cases)

	cnt := len(cases) / 3
	for i := 0; i < cnt; i++ {
		x := sliceToIntIfFloat64(cases[3*i].([]any))
		k := int(cases[3*i+1].(float64))
		want := int64(cases[3*i+2].(float64))
		testEach_kthLargestLevelSum(t, kthLargestLevelSum, x, k, want)
	}
}

func testEach_kthLargestLevelSum(t *testing.T, fn func(*TreeNode, int) int64, x []any, k int, want int64) {
	root := sliceToBinaryTree(x)
	ans := fn(root, k)
	AssertEqual(t, ans, want, fmt.Sprintf("x=%v, k=%v", x, k))
}
