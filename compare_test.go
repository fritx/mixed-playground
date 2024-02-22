package main

import (
	"reflect"
	"testing"
)

func TestDeepEqualExt(t *testing.T) {
	var a []int
	cases := []any{
		[]int{1, 2, 3}, []int{1, 2, 3}, true,
		[]int{1, 2, 3}, []int{1, 2}, false,
		[]int{}, []int{}, true,
		[]int{}, make([]int, 0), true,
		[]int{}, *new([]int), true,
		[]int{}, a, true,
		[]int{}, nil, false,
		[]int{}, []float64{}, false,
		[]int{}, []any{}, false,
	}
	cnt := len(cases) / 3
	for i := 0; i < cnt; i++ {
		x := cases[3*i]
		y := cases[3*i+1]
		want := cases[3*i+2]
		got := deepEqualExt(x, y)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Not equal. x=%v, y=%v, got: %v, want: %v\n", x, y, got, want)
		}
	}
}
