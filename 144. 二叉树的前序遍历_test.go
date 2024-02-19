package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPreorder(t *testing.T) {
	cases := [][]interface{}{}
	json.Unmarshal([]byte(`[
		[1,null,2,3], [1,2,3],
		[], [],
		[1], [1]
	]`), &cases)

	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := sliceToIntIfFloat64(cases[2*i])
		want := toIntSlice(cases[2*i+1])

		root := sliceToBinaryTree(input)
		ans := preorderTraversal(root)

		if !reflect.DeepEqual(ans, want) {
			t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
		}
	}
}
