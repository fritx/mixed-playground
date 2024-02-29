package main

import (
	"encoding/json"
	"playground/utils"
	"reflect"
	"testing"
)

func TestInorder(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,null,2,3], [1,3,2],
		[], [],
		[1], [1]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := sliceToIntIfFloat64(cases[2*i])
		want := utils.ToIntSlice(cases[2*i+1])

		root := sliceToBinaryTree(input)
		ans := inorderTraversal(root)

		if !reflect.DeepEqual(ans, want) {
			t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
		}
	}
}
