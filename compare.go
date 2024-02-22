package main

import (
	"reflect"
)

// 这里拓展reflect.DeepEqual 兼容边界情况下的空切片比较 使返回结果符合预期
func deepEqualExt(x any, y any) bool {
	if isBothEmptySliceLike(x, y) {
		return true
	}
	return reflect.DeepEqual(x, y)
}

// func isBothEmptySliceLike(x any, y any) bool {
// 	if reflect.TypeOf(x) == reflect.TypeOf(y) {
// 		if v := x.([]any); len(v) == 0 {
// 			if v := y.([]any); len(v) == 0 {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
// func isBothEmptySliceLike(x interface{}, y interface{}) bool {
// 	if reflect.TypeOf(x) == reflect.TypeOf(y) {
// 		switch v := x.(type) {
// 		case []interface{}:
// 			if len(v) == 0 {
// 				switch v := y.(type) {
// 				case []interface{}:
// 					if len(v) == 0 {
// 						return true
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

func isBothEmptySliceLike(x interface{}, y interface{}) bool {
	if reflect.TypeOf(x) == reflect.TypeOf(y) {
		return isEmptySlice(x) && isEmptySlice(y)
	}
	return false
}
func isEmptySlice(i interface{}) bool {
	v := reflect.ValueOf(i)
	return v.Kind() == reflect.Slice && v.Len() == 0
}
