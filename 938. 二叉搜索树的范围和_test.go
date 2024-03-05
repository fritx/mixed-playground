package main

import (
	"encoding/json"
	"fmt"
	"playground/utils"
	"testing"
)

func Test_rangeSumBST(t *testing.T) {
	cases := []any{}
	if err := json.Unmarshal([]byte(`[
		[10,5,15,3,7,null,18], 7, 15, 32,
		[10,5,15,3,7,13,18,1,null,6], 6, 10, 23,
		[1, 2], 1, 2, 3,
		[1], 1, 2, 1,
		[1], 2, 3, 0,
		[], 1, 2, 0
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	us := 4
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		data := utils.SliceToIntIfFloat64(cases[us*i].([]any))
		low := int(cases[us*i+1].(float64))
		high := int(cases[us*i+2].(float64))
		want := int(cases[us*i+3].(float64))
		testEach_rangeSumBST(t, rangeSumBST, data, low, high, want)
		testEach_rangeSumBST(t, rangeSumBST_it, data, low, high, want)
	}
}

func testEach_rangeSumBST(t *testing.T, fn func(root *TreeNode, low int, high int) int, data []any, low, high, want int) {
	root := sliceToBinaryTree(data)
	ans := fn(root, low, high)
	utils.AssertEqual(t, ans, want, fmt.Sprintf("data=%v, low=%v, high=%v", data, low, high))
}
