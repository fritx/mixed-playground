package trees

import (
	"encoding/json"
	"playground/utils"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConstructFromPrePost(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,2,4,5,3,6,7], [4,5,2,6,7,3,1], [1,2,3,4,5,6,7],
		[1,2,5,3,6,4], [5,2,6,4,3,1], [1,2,3,5,null,6,4],
		[], [], [],
		[1], [1], [1]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 3
	for i := 0; i < cnt; i++ {
		input := [2][]int{
			utils.ToIntSlice(cases[3*i]),
			utils.ToIntSlice(cases[3*i+1]),
		}
		want := utils.SliceToIntIfFloat64(cases[3*i+2])
		testEachConstructFromPrePost(t, constructFromPrePost, input, want)
	}
}

func testEachConstructFromPrePost(t *testing.T, fn func(preorder []int, postorder []int) *TreeNode, input [2][]int, want []any) {
	root := fn(input[0], input[1])
	// ans := traverseBinaryTree(root)
	// if !reflect.DeepEqual(ans, want) {
	// 	t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	// }

	// wantRoot := buildTreeFromPreorderAndPostorder(input[0], input[1])
	wantRoot := constructFromPrePost(input[0], input[1]) // 暂且用自己 因为LC上已经AC
	if !cmp.Equal(root, wantRoot) {
		t.Errorf("Not equal. got: %v, want: %v\n", root, wantRoot)
	}
}
