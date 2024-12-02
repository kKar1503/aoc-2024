package helper

func SplitArray[T any](arr []T, y int) [][]T {
	n := len(arr)

	// Check if y is greater than the length of the array
	if y > n {
		return [][]T{arr}
	}

	// Calculate the size of each subarray
	subarraySize := n / y
	remainder := n % y

	result := make([][]T, 0, y)
	start := 0

	for i := 0; i < y; i++ {
		end := start + subarraySize
		if i < remainder {
			end++
		}

		result = append(result, arr[start:end])
		start = end
	}

	return result
}

// Filter returns a new slice containing only the elements that satisfy the predicate f.
func Filter[T any](arr []T, f func(T, []T, int) bool) []T {
	result := make([]T, 0, len(arr))
	for i, item := range arr {
		if f(item, result, i) {
			result = append(result, item)
		}
	}
	return result
}

// Map returns a new slice containing the results of applying the function f to each element of the slice.
func Map[T, U any](arr []T, f func(T, []U, int) U) []U {
	result := make([]U, len(arr))
	for i, item := range arr {
		result[i] = f(item, result, i)
	}
	return result
}

func FilterMap[T, U any](arr []T, f func(T, []U, int) (U, bool)) []U {
	result := make([]U, 0, len(arr))
	for i, item := range arr {
		if u, ok := f(item, result, i); ok {
			result = append(result, u)
		}
	}
	return result
}

func Contains[T comparable](arr []T, item T) bool {
	for _, i := range arr {
		if i == item {
			return true
		}
	}

	return false
}

func ContainsFunc[T any](arr []T, f func(T) bool) bool {
	for _, i := range arr {
		if f(i) {
			return true
		}
	}

	return false
}

func GroupBy[T any, K comparable](arr []T, f func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, item := range arr {
		k := f(item)
		m[k] = append(m[k], item)
	}
	return m
}

func MapValues[T comparable, U any](m map[T]U) []U {
	result := make([]U, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func RemoveByIndex[T any](arr []T, i int) []T {
	clone := make([]T, len(arr))
	copy(clone, arr)
	return append(clone[:i], clone[i+1:]...)
}

func Some[T any](arr []T, f func(T) bool) bool {
	for _, item := range arr {
		if f(item) {
			return true
		}
	}
	return false
}
