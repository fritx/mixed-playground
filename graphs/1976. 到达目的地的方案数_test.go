package graphs

import (
	"encoding/json"
	"fmt"
	"playground/utils"
	"testing"
)

func Test_countPaths(t *testing.T) {
	cases := []any{}
	if err := json.Unmarshal([]byte(`[
		7, [[0,6,7],[0,1,2],[1,2,3],[1,3,3],[6,3,3],[3,5,1],[6,5,1],[2,5,1],[0,4,5],[4,6,2]], 4,
        2, [[1,0,10]], 1,
        2, [], 0
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	us := 3
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		n := int(cases[us*i].(float64))

		// fix: panic: interface conversion: interface {} is []interface {}, not [][]interface {}
		// _roads := cases[us*i+1].([][]any)
		_roads := utils.SliceItemTypeAssert[[]any](cases[us*i+1].([]any))

		roads := make([][]int, len(_roads))
		for i, r := range _roads {
			roads[i] = utils.ToIntSlice(r)
		}
		want := int(cases[us*i+2].(float64))
		testEach_countPaths(t, countPaths, n, roads, want)
	}
}

func testEach_countPaths(t *testing.T, fn func(int, [][]int) int, n int, roads [][]int, want int) {
	ans := fn(n, roads)
	utils.AssertEqual(t, ans, want, fmt.Sprintf("n=%v, roads=%v", n, roads))
}
