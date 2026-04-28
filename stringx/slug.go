package stringx

import (
	"regexp"
	"strings"
)

var (
	regexpNonAlphanumeric = regexp.MustCompile(`[^a-z0-9]+`)
	regexpMultipleHyphens = regexp.MustCompile(`-+`)
)

// Slug converts a string into a URL-friendly slug.
// It converts characters to lowercase, replaces spaces and other non-alphanumeric
// characters with hyphens, and removes redundant hyphens.
func Slug(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Simple transliteration for Polish characters
	replacer := strings.NewReplacer(
		"ą", "a", "ć", "c", "ę", "e", "ł", "l", "ń", "n", "ó", "o", "ś", "s", "ź", "z", "ż", "z",
	)
	s = replacer.Replace(s)

	// Replace all non-alphanumeric characters with hyphens
	s = regexpNonAlphanumeric.ReplaceAllString(s, "-")

	// Replace multiple hyphens with a single one
	s = regexpMultipleHyphens.ReplaceAllString(s, "-")

	// Trim hyphens from ends
	return strings.Trim(s, "-")
}
