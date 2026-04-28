package slicex

import (
	"reflect"
	"testing"
	"time"
)

func TestSortStrings(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "unsorted strings",
			input:    []string{"banana", "apple", "cherry"},
			expected: []string{"apple", "banana", "cherry"},
		},
		{
			name:     "already sorted strings",
			input:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortStrings(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SortStrings() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestSortDates(t *testing.T) {
	d1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	d3 := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    []time.Time
		expected []time.Time
	}{
		{
			name:     "unsorted dates",
			input:    []time.Time{d1, d2, d3},
			expected: []time.Time{d2, d1, d3},
		},
		{
			name:     "already sorted dates",
			input:    []time.Time{d2, d1, d3},
			expected: []time.Time{d2, d1, d3},
		},
		{
			name:     "empty slice",
			input:    []time.Time{},
			expected: []time.Time{},
		},
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SortDates(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("SortDates() = %v, want %v", got, tt.expected)
			}
		})
	}
}
