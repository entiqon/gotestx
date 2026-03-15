package gotestx

import (
	"runtime"
	"testing"
)

func TestGetGOOS(t *testing.T) {
	t.Run("Default", func(t *testing.T) {
		if getGOOS() != runtime.GOOS {
			t.Fatalf("expected %s, got %s", runtime.GOOS, getGOOS())
		}
	})

	t.Run("Override", func(t *testing.T) {
		original := getGOOS
		defer func() { getGOOS = original }()

		getGOOS = func() string {
			return "test-os"
		}

		if getGOOS() != "test-os" {
			t.Fatalf("expected overridden OS")
		}
	})
}
