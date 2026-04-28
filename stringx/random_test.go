package stringx

import (
	"strings"
	"testing"
)

func TestRandomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"zero length", 0},
		{"negative length", -1},
		{"positive length 1", 1},
		{"positive length 10", 10},
		{"positive length 100", 100},
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.length)
			if tt.length <= 0 {
				if got != "" {
					t.Errorf("RandomString(%d) = %q, want empty string", tt.length, got)
				}

				return
			}

			if len(got) != tt.length {
				t.Errorf("RandomString(%d) length = %d, want %d", tt.length, len(got), tt.length)
			}

			for _, char := range got {
				if !strings.ContainsRune(charset, char) {
					t.Errorf("RandomString(%d) contains invalid character %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandomSmallString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"zero length", 0},
		{"positive length 10", 10},
	}

	charset := "abcdefghijklmnopqrstuvwxyz"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomSmallString(tt.length)
			if tt.length <= 0 {
				if got != "" {
					t.Errorf("RandomSmallString(%d) = %q, want empty string", tt.length, got)
				}

				return
			}

			if len(got) != tt.length {
				t.Errorf("RandomSmallString(%d) length = %d, want %d", tt.length, len(got), tt.length)
			}

			for _, char := range got {
				if !strings.ContainsRune(charset, char) {
					t.Errorf("RandomSmallString(%d) contains invalid character %q", tt.length, char)
				}
			}
		})
	}
}

func TestRandomString_Randomness(t *testing.T) {
	s1 := RandomString(32)
	s2 := RandomString(32)

	if s1 == s2 {
		t.Errorf("RandomString(32) generated two identical strings: %q", s1)
	}
}
