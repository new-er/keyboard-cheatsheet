package main

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Flatten[T any, U any](slice []T, fn func(T) []U) []U {
	var result []U
	for _, v := range slice {
		result = append(result, fn(v)...)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Count[T any](slice []T, fn func(T) bool) int {
	count := 0
	for _, v := range slice {
		if fn(v) {
			count++
		}
	}
	return count
}
