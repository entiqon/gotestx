package gotestx

import "testing"

func TestCoverage(t *testing.T) {
	t.Run("BuildCoverage", func(t *testing.T) {
		t.Run("Args", func(t *testing.T) {
			pkgs := []string{"./...", "./internal"}

			args := buildCoverageArgs(pkgs)

			expectedPrefix := []string{
				"test",
				"-coverprofile=" + coverageFile,
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
			cmd := buildCoverageOpenCommand()

			if cmd == nil {
				t.Fatalf("expected command instance")
			}
		})
	})

}
