package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	cases := [][]interface{}{}
	json.Unmarshal([]byte(`[
		[1,null,2,3], [3,2,1],
		[], [],
		[1], [1]
	]`), &cases)

	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := sliceToIntIfFloat64(cases[2*i])
		want := toIntSlice(cases[2*i+1])

		testFn(t, postorderTraversal, input, want)
		testFn(t, postorderTraversal_2, input, want)
	}
}

func testFn(t *testing.T, fn func(*TreeNode) []int, input []interface{}, want []int) {
	root := sliceToBinaryTree(input)
	ans := fn(root)

	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	}
}
