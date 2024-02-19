package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPostorderNTree(t *testing.T) {
	cases := [][]interface{}{}
	json.Unmarshal([]byte(`[
		[1,null,3,2,4,null,5,6], [5,6,3,2,4,1],
		[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14],
			[2,6,14,11,7,3,12,8,4,13,9,10,5,1],
		[], [],
		[1], [1]
	]`), &cases)

	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := sliceToIntIfFloat64(cases[2*i])
		want := toIntSlice(cases[2*i+1])

		testNTreeFn(t, postorder, input, want)
		// testNTreeFn(t, postorder_it, input, want)
	}
}

func testNTreeFn(t *testing.T, fn func(*Node) []int, input []interface{}, want []int) {
	// root := sliceToNTree(input)
	root := convertSliceToNaryTree(input)
	ans := fn(root)

	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	}
}
