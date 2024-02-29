package main

func Map[T, U any](arr []T, transform func(T) U) []U {
	result := make([]U, len(arr))
	for i, v := range arr {
		result[i] = transform(v)
	}
	return result
}
