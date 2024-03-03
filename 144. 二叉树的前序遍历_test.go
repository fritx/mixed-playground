package main

import (
	"encoding/json"
	"playground/utils"
	"reflect"
	"testing"
)

func TestPreorder(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,null,2,3], [1,2,3],
		[], [],
		[1], [1]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := utils.SliceToIntIfFloat64(cases[2*i])
		want := utils.ToIntSlice(cases[2*i+1])

		testEachPreorder(t, preorderTraversal, input, want)
		testEachPreorder(t, preorderTraversal_2, input, want)
		testEachPreorder(t, preorderTraversal_dfs, input, want)
	}
}

func testEachPreorder(t *testing.T, fn func(*TreeNode) []int, input []any, want []int) {
	root := sliceToBinaryTree(input)
	ans := fn(root)

	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. data=%v. got: %v, want: %v\n", input, ans, want)
	}
}
