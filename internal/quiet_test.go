package gotestx_test

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	gotestx "github.com/entiqon/gotestx/internal"
)

func TestQuiet(t *testing.T) {
	t.Run("Failure", func(t *testing.T) {
		var out bytes.Buffer
		var errBuf bytes.Buffer

		opts := &gotestx.Options{}

		code := gotestx.HandleQuietOutput(
			"",
			errors.New("test failure"),
			opts,
			&out,
			&errBuf,
		)

		if code != 1 {
			t.Fatalf("expected exit code 1, got %d", code)
		}

		if !strings.Contains(errBuf.String(), "Tests failed") {
			t.Fatalf("expected failure message")
		}
	})

	t.Run("Success", func(t *testing.T) {
		t.Run("WithoutCoverage", func(t *testing.T) {
			var out bytes.Buffer
			var errBuf bytes.Buffer

			opts := &gotestx.Options{}

			code := gotestx.HandleQuietOutput(
				"ok example/pkg 0.001s",
				nil,
				opts,
				&out,
				&errBuf,
			)

			if code != gotestx.ExitOK {
				t.Fatalf("expected ExitOK, got %d", code)
			}

			if !strings.Contains(out.String(), "Tests finished successfully") {
				t.Fatalf("expected success message")
			}
		})
	})

	t.Run("CoverageSummary", func(t *testing.T) {

		var out bytes.Buffer
		var errBuf bytes.Buffer

		opts := &gotestx.Options{
			WithCoverage: true,
		}

		buf := `
ok   github.com/example/a  0.01s
ok   github.com/example/b  0.02s
coverage: 92.3% of statements
`

		code := gotestx.HandleQuietOutput(
			buf,
			nil,
			opts,
			&out,
			&errBuf,
		)

		if code != gotestx.ExitOK {
			t.Fatalf("expected ExitOK")
		}

		if !strings.Contains(out.String(), "coverage:") {
			t.Fatalf("expected coverage line in output")
		}
	})
}
