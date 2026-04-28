package slicex

import (
	"slices"
	"time"
)

// SortStrings sorts a slice of strings in ascending order and returns it.
// It modifies the original slice.
func SortStrings(s []string) []string {
	slices.Sort(s)

	return s
}

// SortDates sorts a slice of time.Time in ascending order and returns it.
// It modifies the original slice.
func SortDates(s []time.Time) []time.Time {
	slices.SortFunc(s, func(a, b time.Time) int {
		return a.Compare(b)
	})

	return s
}
