package main

func Filter[T any](arr []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(arr))
	for _, v := range arr {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
