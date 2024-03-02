package linked_lists

import (
	"encoding/json"
	"playground/utils"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	cases := [][]any{}
	if err := json.Unmarshal([]byte(`[
		[1,2,3,4,5], [5,4,3,2,1],
		[1,2], [2,1],
		[1], [1],
		[], []
	]`), &cases); err != nil {
		t.Fatalf("json.Unmarshal error: %v\n", err)
	}
	cnt := len(cases) / 2
	for i := 0; i < cnt; i++ {
		data := utils.ToIntSlice(cases[2*i])
		want := utils.ToIntSlice(cases[2*i+1])

		testEach_reverseList(t, reverseList, data, want)
		testEach_reverseList(t, reverseList_it, data, want)
	}
}

func testEach_reverseList(t *testing.T, fn func(*ListNode) *ListNode, data []int, want []int) {
	head := sliceToLinkedList(data)
	ans := linkedListToSlice(fn(head))

	if !reflect.DeepEqual(ans, want) {
		t.Errorf("Not equal. got: %v, want: %v\n", ans, want)
	}
}
