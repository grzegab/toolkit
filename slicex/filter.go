package slicex

// Filter returns a new slice containing all elements of s that satisfy the predicate keep.
func Filter[T any](s []T, keep func(T) bool) []T {
	var res []T

	for _, v := range s {
		if keep(v) {
			res = append(res, v)
		}
	}

	return res
}

// Intersect returns a new slice containing elements that are present in both s1 and s2.
func Intersect[T comparable](s1, s2 []T) []T {
	set := make(map[T]struct{})
	for _, v := range s2 {
		set[v] = struct{}{}
	}

	var res []T

	for _, v := range s1 {
		if _, ok := set[v]; ok {
			res = append(res, v)
		}
	}

	return res
}

// Difference returns a new slice containing elements from s1 that are not present in s2.
func Difference[T comparable](s1, s2 []T) []T {
	set := make(map[T]struct{})
	for _, v := range s2 {
		set[v] = struct{}{}
	}

	var res []T

	for _, v := range s1 {
		if _, ok := set[v]; !ok {
			res = append(res, v)
		}
	}

	return res
}

// RemoveZero returns a new slice with all zero values of type T removed from s.
func RemoveZero[T comparable](s []T) []T {
	var zero T

	return Filter(s, func(v T) bool {
		return v != zero
	})
}

// Unique returns a new slice with all duplicate elements removed from s, preserving the original order.
func Unique[T comparable](s []T) []T {
	seen := make(map[T]struct{})

	var res []T

	for _, v := range s {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			res = append(res, v)
		}
	}

	return res
}
