package datex

import "time"

// Parse date string to time.Time
func Parse(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

// RemoveHours removes hours, minutes, seconds and nanoseconds from time
func RemoveHours(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

// RoundToHour rounds time to nearest hour
func RoundToHour(t time.Time) time.Time {
	return t.Round(time.Hour)
}

// RoundToDay rounds time to nearest day
func RoundToDay(t time.Time) time.Time {
	return t.Round(24 * time.Hour)
}
