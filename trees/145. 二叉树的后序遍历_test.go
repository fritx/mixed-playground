package trees

import (
	"encoding/json"
	"playground/utils"
	"reflect"
	"testing"
)

func TestPostorder(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,null,2,3], [3,2,1],
		[], [],
		[1], [1]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		input := utils.SliceToIntIfFloat64(cases[2*i])
		want := utils.ToIntSlice(cases[2*i+1])

		testEachPostorder(t, postorderTraversal, input, want)
		testEachPostorder(t, postorderTraversal_2, input, want)
	}
}

func testEachPostorder(t *testing.T, fn func(*TreeNode) []int, input []any, want []int) {
	root := sliceToBinaryTree(input)
	ans := fn(root)

	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	}
}
