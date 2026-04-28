package slicex

import (
	"reflect"
	"testing"
)

func TestIntsToStrings(t *testing.T) {
	tests := []struct {
		name string
		ints []int
		want []string
	}{
		{
			name: "multiple integers",
			ints: []int{1, 2, 3, 4, 5},
			want: []string{"1", "2", "3", "4", "5"},
		},
		{
			name: "negative and zero",
			ints: []int{-1, 0, 1},
			want: []string{"-1", "0", "1"},
		},
		{
			name: "empty slice",
			ints: []int{},
			want: []string{},
		},
		{
			name: "nil slice",
			ints: nil,
			want: []string{},
		},
		{
			name: "large integers",
			ints: []int{123456789, -987654321},
			want: []string{"123456789", "-987654321"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntsToStrings(tt.ints); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntsToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
