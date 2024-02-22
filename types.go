package main

import "fmt"

func toIntSlice(slice []any) []int {
	res := make([]int, len(slice))
	for i, value := range slice {
		switch v := value.(type) {
		case int:
			res[i] = v
		case float64:
			res[i] = int(v)
		}
	}
	return res
}

func sliceToIntIfFloat64(slice []any) []any {
	res := make([]any, len(slice))
	for i, v := range slice {
		iv, _ := convertToFloat64ToInt(v)
		res[i] = iv
	}
	return res
}

func convertToFloat64ToInt(value any) (any, error) {
	switch v := value.(type) {
	case float64:
		// 将float64转换为int，注意这可能会丢失精度或导致溢出
		intValue := int(v)
		// 检查是否溢出
		if float64(intValue) != v {
			return 0, fmt.Errorf("float64 value %f overflows int", v)
		}
		return intValue, nil
	case int:
		return v, nil
	case nil:
		return v, nil
	default:
		// 如果不是float64，返回错误
		return 0, fmt.Errorf("unsupported type %T for conversion to int", value)
	}
}
