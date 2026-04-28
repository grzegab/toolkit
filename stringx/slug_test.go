package stringx

import "testing"

func TestSlug(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple text",
			input:    "hello world",
			expected: "hello-world",
		},
		{
			name:     "uppercase",
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			name:     "multiple spaces",
			input:    "hello   world",
			expected: "hello-world",
		},
		{
			name:     "special characters",
			input:    "hello world!",
			expected: "hello-world",
		},
		{
			name:     "numbers and hyphens",
			input:    "version 1.2.3 released",
			expected: "version-1-2-3-released",
		},
		{
			name:     "leading and trailing spaces",
			input:    "  hello world  ",
			expected: "hello-world",
		},
		{
			name:     "non-ascii characters",
			input:    "zażółć gęślą jaźń",
			expected: "zazolc-gesla-jazn",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only special characters",
			input:    "!!!@@@###",
			expected: "",
		},
		{
			name:     "complex string",
			input:    "  This is a TEST... with some SPECIAL characters! 123 ",
			expected: "this-is-a-test-with-some-special-characters-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Slug(tt.input)
			if actual != tt.expected {
				t.Errorf("Slug(%q) = %q; want %q", tt.input, actual, tt.expected)
			}
		})
	}
}
