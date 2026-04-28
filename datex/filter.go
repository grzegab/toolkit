package datex

import (
	"errors"
	"time"
)

var ErrInvalidFormat = errors.New("invalid date format")

type DateConstraint interface {
	~string | time.Time
}

type Direction int

const (
	Older Direction = iota
	Newer
)

// Filter filters dates relative to the base date.
// If direction is Newer, it returns dates newer than or equal to base (inclusive).
// If direction is Older, it returns dates older than or equal to base (inclusive).
func Filter[T DateConstraint](base T, dates []T, direction Direction) []T {
	baseTime, err := toTime(base)
	if err != nil {
		return nil
	}

	var result []T

	for _, d := range dates {
		t, err := toTime(d)
		if err != nil {
			continue
		}

		switch direction {
		case Newer:
			if t.After(baseTime) || t.Equal(baseTime) {
				result = append(result, d)
			}
		case Older:
			if t.Before(baseTime) || t.Equal(baseTime) {
				result = append(result, d)
			}
		}
	}

	return result
}

func toTime(v any) (time.Time, error) {
	switch t := v.(type) {
	case time.Time:
		return t, nil
	case string:
		// Try various popular formats
		formats := []string{
			time.RFC3339,
			"2006-01-02 15:04:05",
			"2006-01-02",
		}
		for _, f := range formats {
			if parsed, err := time.Parse(f, t); err == nil {
				return parsed, nil
			}
		}
	}

	return time.Time{}, ErrInvalidFormat
}
