package gotestx_test

import (
	"testing"

	gotestx "github.com/entiqon/gotestx/internal"
)

func TestCoverage(t *testing.T) {
	t.Run("BuildCoverage", func(t *testing.T) {
		t.Run("Args", func(t *testing.T) {
			pkgs := []string{"./...", "./internal"}

			args := gotestx.BuildCoverageArgs(pkgs)

			expectedPrefix := []string{
				"test",
				"-coverprofile=" + gotestx.CoverageFile,
				"-covermode=atomic",
			}

			for i, v := range expectedPrefix {
				if args[i] != v {
					t.Fatalf("expected %s at position %d, got %s", v, i, args[i])
				}
			}

			if len(args) != len(expectedPrefix)+len(pkgs) {
				t.Fatalf("unexpected args length")
			}
		})

		t.Run("OpenCommand", func(t *testing.T) {
			cmd := gotestx.BuildCoverageOpenCommand()

			if cmd == nil {
				t.Fatalf("expected command instance")
			}

		})
	})

}
