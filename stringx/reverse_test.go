package stringx

import (
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "empty string",
			s:    "",
			want: "",
		},
		{
			name: "single character",
			s:    "a",
			want: "a",
		},
		{
			name: "ascii string",
			s:    "hello",
			want: "olleh",
		},
		{
			name: "unicode string (polish)",
			s:    "zażółć gęślą jaźń",
			want: "ńźaj ąlśęg ćłóżaz",
		},
		{
			name: "unicode string (emoji)",
			s:    "hello 👋 world 🌍",
			want: "🌍 dlrow 👋 olleh",
		},
		{
			name: "palindrome",
			s:    "kajak",
			want: "kajak",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.s); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
