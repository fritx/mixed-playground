package utils

import "fmt"

func SliceToIntIfFloat64(slice []any) []any {
	res := make([]any, len(slice))
	for i, v := range slice {
		iv, err := convertToIntIfFloat64(v)
		if err != nil {
			panic(err) // todo: don't panic
		}
		res[i] = iv
	}
	return res
}

// allowing more json value types
func convertToIntIfFloat64(value any) (any, error) {
	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		return v, nil
	}
	return convertToIntIfFloat64OnlyIntAndNil(value)
}

// allowing only null and numbers
func convertToIntIfFloat64OnlyIntAndNil(value any) (any, error) {
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
