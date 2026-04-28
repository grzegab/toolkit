package grpcx

import (
	"testing"
)

func TestConvertInt(t *testing.T) {
	i32 := int32(20)
	i64 := int64(30)
	i := 10

	tests := []struct {
		input    any
		name     string
		expected int
	}{
		{nil, "nil", 0},
		{10, "int", 10},
		{int32(20), "int32", 20},
		{int64(30), "int64", 30},
		{float32(40.5), "float32", 40},
		{float64(50.9), "float64", 50},
		{&i, "*int", 10},
		{&i32, "*int32", 20},
		{&i64, "*int64", 30},
		{(*int)(nil), "*int nil", 0},
		{"100", "string (unsupported)", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertInt(tt.input); got != tt.expected {
				t.Errorf("ConvertInt() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertFloat(t *testing.T) {
	f32 := float32(10.5)
	f64 := float64(20.5)

	tests := []struct {
		input    any
		name     string
		expected float64
	}{
		{nil, "nil", 0.0},
		{float32(10.5), "float32", 10.5},
		{float64(20.5), "float64", 20.5},
		{30, "int", 30.0},
		{int32(40), "int32", 40.0},
		{int64(50), "int64", 50.0},
		{&f32, "*float32", 10.5},
		{&f64, "*float64", 20.5},
		{(*float64)(nil), "*float64 nil", 0.0},
		{"100", "string (unsupported)", 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertFloat(tt.input); got != tt.expected {
				t.Errorf("ConvertFloat() = %v, want %v", got, tt.expected)
			}
		})
	}
}
