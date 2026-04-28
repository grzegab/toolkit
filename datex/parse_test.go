package datex

import (
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	input := "2026-04-27T12:00:00Z"

	got, err := Parse(input)
	if err != nil {
		t.Fatalf("Parse(%q) error: %v", input, err)
	}

	want := time.Date(2026, 4, 27, 12, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("Parse(%q) = %v, want %v", input, got, want)
	}
}

func TestRemoveHours(t *testing.T) {
	tm := time.Date(2026, 4, 27, 12, 30, 45, 100, time.UTC)
	got := RemoveHours(tm)

	want := time.Date(2026, 4, 27, 0, 0, 0, 0, time.UTC)
	if !got.Equal(want) {
		t.Errorf("RemoveHours(%v) = %v, want %v", tm, got, want)
	}
}

func TestRoundToHour(t *testing.T) {
	tests := []struct {
		input time.Time
		want  time.Time
	}{
		{
			input: time.Date(2026, 4, 27, 12, 29, 59, 0, time.UTC),
			want:  time.Date(2026, 4, 27, 12, 0, 0, 0, time.UTC),
		},
		{
			input: time.Date(2026, 4, 27, 12, 30, 0, 0, time.UTC),
			want:  time.Date(2026, 4, 27, 13, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		got := RoundToHour(tt.input)
		if !got.Equal(tt.want) {
			t.Errorf("RoundToHour(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestRoundToDay(t *testing.T) {
	tests := []struct {
		input time.Time
		want  time.Time
	}{
		{
			input: time.Date(2026, 4, 27, 11, 59, 59, 0, time.UTC),
			want:  time.Date(2026, 4, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			input: time.Date(2026, 4, 27, 12, 0, 0, 0, time.UTC),
			want:  time.Date(2026, 4, 28, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		got := RoundToDay(tt.input)
		if !got.Equal(tt.want) {
			t.Errorf("RoundToDay(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
