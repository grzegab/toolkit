package datex

import (
	"reflect"
	"testing"
	"time"
)

const baseDate = "2023-01-01"

func TestFilter(t *testing.T) {
	now := time.Now()
	past := now.Add(-24 * time.Hour)
	future := now.Add(24 * time.Hour)

	t.Run("time.Time newer", func(t *testing.T) {
		dates := []time.Time{past, now, future}
		got := Filter(now, dates, Newer)

		want := []time.Time{now, future}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("time.Time older", func(t *testing.T) {
		dates := []time.Time{past, now, future}
		got := Filter(now, dates, Older)

		want := []time.Time{past, now}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("string newer", func(t *testing.T) {
		base := baseDate
		dates := []string{"2022-12-31", baseDate, "2023-01-02"}
		got := Filter(base, dates, Newer)

		want := []string{baseDate, "2023-01-02"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("string older", func(t *testing.T) {
		base := baseDate
		dates := []string{"2022-12-31", baseDate, "2023-01-02"}
		got := Filter(base, dates, Older)

		want := []string{"2022-12-31", baseDate}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("mixed string formats", func(t *testing.T) {
		base := baseDate
		dates := []string{"2022-12-31 10:00:00", "2023-01-01T12:00:00Z", "2023-01-02"}
		got := Filter(base, dates, Newer)

		want := []string{"2023-01-01T12:00:00Z", "2023-01-02"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
