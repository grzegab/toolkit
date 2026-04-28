package uuidx

import (
	"testing"

	"github.com/google/uuid"
)

func TestGenerators(t *testing.T) {
	tests := []struct {
		fn   func() string
		name string
	}{
		{V1, "V1"},
		{V2, "V2"},
		{V4, "V4"},
		{V6, "V6"},
		{V7, "V7"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fn()
			if _, err := uuid.Parse(got); err != nil {
				t.Errorf("%s() generated invalid UUID: %q, error: %v", tt.name, got, err)
			}
		})
	}
}

func TestV3V5(t *testing.T) {
	namespaces := []struct {
		name string
		val  string
	}{
		{"DNS", NamespaceDNS},
		{"URL", NamespaceURL},
		{"OID", NamespaceOID},
		{"X500", NamespaceX500},
	}
	name := "example.com"

	for _, ns := range namespaces {
		t.Run("V3_"+ns.name, func(t *testing.T) {
			got := V3(ns.val, name)

			u, err := uuid.Parse(got)
			if err != nil {
				t.Errorf("V3(%s) generated invalid UUID: %q, error: %v", ns.name, got, err)
			}

			if u.Version() != 3 {
				t.Errorf("V3(%s) version = %v, want 3", ns.name, u.Version())
			}
		})

		t.Run("V5_"+ns.name, func(t *testing.T) {
			got := V5(ns.val, name)

			u, err := uuid.Parse(got)
			if err != nil {
				t.Errorf("V5(%s) generated invalid UUID: %q, error: %v", ns.name, got, err)
			}

			if u.Version() != 5 {
				t.Errorf("V5(%s) version = %v, want 5", ns.name, u.Version())
			}
		})
	}

	t.Run("V3InvalidNS", func(t *testing.T) {
		got := V3("invalid", name)

		if got != uuid.Nil.String() {
			t.Errorf("V3() with invalid NS = %q, want %q", got, uuid.Nil.String())
		}
	})
}

func TestVersions(t *testing.T) {
	checkVersion := func(t *testing.T, name string, got string, want uuid.Version) {
		u, err := uuid.Parse(got)
		if err != nil {
			t.Fatalf("%s: failed to parse %q: %v", name, got, err)
		}

		if u.Version() != want {
			t.Errorf("%s: version = %v, want %v", name, u.Version(), want)
		}
	}

	checkVersion(t, "V1", V1(), 1)
	checkVersion(t, "V2", V2(), 2)
	checkVersion(t, "V4", V4(), 4)
	checkVersion(t, "V6", V6(), 6)
	checkVersion(t, "V7", V7(), 7)
}
