package data_structures

import (
	"encoding/json"
	"playground/utils"
	"reflect"
	"testing"
)

func Test_stackByQueues(t *testing.T) {
	cases := []any{}
	if err := json.Unmarshal([]byte(`[
		["MyStack","push","push","top","pop","empty"], [[],[1],[2],[],[],[]], [null,null,null,2,2,false],
		["MyStack"], [[]], [null]
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	us := 3
	cnt := len(cases) / us
	for i := 0; i < cnt; i++ {
		// panic: interface conversion: interface {} is []interface {}, not []string
		// ls1 := cases[us*i].([]string)
		ls1 := utils.SliceItemTypeAssert[string](cases[us*i].([]any))

		// panic: interface conversion: interface {} is []interface {}, not [][]interface {}
		// ls2 := cases[us*i+1].([][]any)
		ls2 := utils.SliceItemTypeAssert[[]any](cases[us*i+1].([]any))

		for i, v := range ls2 {
			ls2[i] = utils.SliceToIntIfFloat64(v)
		}
		want := utils.SliceToIntIfFloat64(cases[us*i+2].([]any))

		ls1 = ls1[1:]
		ls2 = ls2[1:]
		want = want[1:]
		s := Constructor()
		s2 := Constructor2()
		s3 := Constructor3()
		testEach_stackByQueues(t, &s, ls1, ls2, want)
		testEach_stackByQueues(t, &s2, ls1, ls2, want)
		testEach_stackByQueues(t, &s3, ls1, ls2, want)
	}
}

func testEach_stackByQueues(t *testing.T, s Stack[int], ls1 []string, ls2 [][]any, want []any) {
	ans := make([]any, len(want))
	for i, me := range ls1 {
		switch me {
		case "push":
			s.Push(ls2[i][0].(int))
			ans[i] = nil
		case "top":
			ans[i] = s.Top()
		case "pop":
			ans[i] = s.Pop()
		case "empty":
			ans[i] = s.Empty()
		}
	}
	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	}
}
