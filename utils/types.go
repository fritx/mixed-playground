package utils

func SliceItemTypeAssert[T any](x []any) []T {
	r := make([]T, len(x))
	for i, v := range x {
		r[i] = v.(T)
	}
	return r
}
