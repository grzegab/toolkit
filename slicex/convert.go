package slicex

import "strconv"

// IntsToStrings converts a slice of integers to a slice of strings.
func IntsToStrings(ints []int) []string {
	res := make([]string, len(ints))
	for i, v := range ints {
		res[i] = strconv.Itoa(v)
	}

	return res
}
