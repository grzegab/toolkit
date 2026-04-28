package slicex

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		keep func(int) bool
		want []int
	}{
		{
			name: "filter even numbers",
			s:    []int{1, 2, 3, 4, 5, 6},
			keep: func(n int) bool { return n%2 == 0 },
			want: []int{2, 4, 6},
		},
		{
			name: "filter greater than 10",
			s:    []int{5, 15, 2, 20},
			keep: func(n int) bool { return n > 10 },
			want: []int{15, 20},
		},
		{
			name: "empty slice",
			s:    []int{},
			keep: func(n int) bool { return n > 0 },
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.s, tt.keep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "common elements",
			s1:   []int{1, 2, 3, 4},
			s2:   []int{3, 4, 5, 6},
			want: []int{3, 4},
		},
		{
			name: "no common elements",
			s1:   []int{1, 2},
			s2:   []int{3, 4},
			want: nil,
		},
		{
			name: "empty s1",
			s1:   []int{},
			s2:   []int{1, 2},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.s1, tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		s1   []int
		s2   []int
		want []int
	}{
		{
			name: "remove elements",
			s1:   []int{1, 2, 3, 4},
			s2:   []int{2, 4},
			want: []int{1, 3},
		},
		{
			name: "remove nothing",
			s1:   []int{1, 2, 3},
			s2:   []int{4, 5},
			want: []int{1, 2, 3},
		},
		{
			name: "remove everything",
			s1:   []int{1, 2},
			s2:   []int{1, 2, 3},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Difference(tt.s1, tt.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Difference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveZero(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "remove zeroes from int slice",
			s:    []int{0, 1, 0, 2, 3, 0},
			want: []int{1, 2, 3},
		},
		{
			name: "no zeroes",
			s:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all zeroes",
			s:    []int{0, 0, 0},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveZero(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveZero() = %v, want %v", got, tt.want)
			}
		})
	}

	t.Run("remove empty strings", func(t *testing.T) {
		s := []string{"", "a", "", "b"}

		want := []string{"a", "b"}
		if got := RemoveZero(s); !reflect.DeepEqual(got, want) {
			t.Errorf("RemoveZero() = %v, want %v", got, want)
		}
	})
}

func TestUnique(t *testing.T) {
	tests := []struct {
		name string
		s    []int
		want []int
	}{
		{
			name: "remove duplicates",
			s:    []int{1, 2, 1, 3, 2, 4},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "no duplicates",
			s:    []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			name: "all duplicates",
			s:    []int{1, 1, 1},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
