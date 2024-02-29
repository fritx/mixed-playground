package main

import (
	"encoding/json"
	"playground/utils"
	"testing"
)

func TestPostorderNTree(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,null,3,2,4,null,5,6], [5,6,3,2,4,1],
		[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14],
			[2,6,14,11,7,3,12,8,4,13,9,10,5,1],
		[], [],
		[1], [1]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := sliceToIntIfFloat64(cases[2*i])
		want := utils.ToIntSlice(cases[2*i+1])

		testEachNTree(t, postorder, input, want)
		testEachNTree(t, postorder_it, input, want)
	}
}

func testEachNTree(t *testing.T, fn func(*Node) []int, input []any, want []int) {
	// root := sliceToNTree(input)
	// root := convertSliceToNaryTree(input)
	root := createTree(input)
	ans := fn(root)

	// 这里拓展reflect.DeepEqual 兼容边界情况下的空切片比较 使返回结果符合预期
	// if !reflect.DeepEqual(ans, want) {
	if !deepEqualExt(ans, want) {
		t.Errorf("Not equal. got: %T-%v, want: %T-%v\n", ans, ans, want, want)
	}
}
