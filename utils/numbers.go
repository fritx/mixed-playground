package utils

func ToIntSlice(slice []any) []int {
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
