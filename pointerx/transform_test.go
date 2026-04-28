package pointerx

import (
	"testing"
)

func TestToPointer(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		val := 42

		ptr := ToPointer(val)
		if ptr == nil {
			t.Fatal("expected pointer to be not nil")
		}

		if *ptr != val {
			t.Errorf("expected %d, got %d", val, *ptr)
		}
		// Sprawdzenie czy to kopia
		val = 100
		if *ptr == val {
			t.Error("expected pointer value to not change when original value changes")
		}
	})

	t.Run("string", func(t *testing.T) {
		val := "hello"

		ptr := ToPointer(val)
		if ptr == nil {
			t.Fatal("expected pointer to be not nil")
		}

		if *ptr != val {
			t.Errorf("expected %s, got %s", val, *ptr)
		}
	})

	t.Run("struct", func(t *testing.T) {
		type testStruct struct {
			B string
			A int
		}

		val := testStruct{A: 1, B: "test"}

		ptr := ToPointer(val)
		if ptr == nil {
			t.Fatal("expected pointer to be not nil")
		}

		if *ptr != val {
			t.Errorf("expected %+v, got %+v", val, *ptr)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		val := 42
		origPtr := &val

		ptr := ToPointer(origPtr)
		if ptr == nil {
			t.Fatal("expected pointer to be not nil")
		}

		if *ptr != origPtr {
			t.Errorf("expected %p, got %p", origPtr, *ptr)
		}

		if **ptr != 42 {
			t.Errorf("expected 42, got %d", **ptr)
		}
	})
}
