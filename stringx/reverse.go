package stringx

import "slices"

// Reverse zwraca odwrócony ciąg znaków (rune po rune).
func Reverse(s string) string {
	r := []rune(s)
	slices.Reverse(r)

	return string(r)
}
