package gotestx_test

import (
	"bytes"
	"io"
	"testing"

	gotestx "github.com/entiqon/gotestx/internal"
)

func TestResolveOptions(t *testing.T) {

	t.Run("Packages", func(t *testing.T) {

		t.Run("Default", func(t *testing.T) {
			opts, code := gotestx.ResolveOptions([]string{}, io.Discard, io.Discard)

			if code != gotestx.ExitContinue {
				t.Fatalf("expected ExitContinue, got %d", code)
			}

			if opts == nil {
				t.Fatalf("expected options not nil")
			}

			if len(opts.Packages) != 1 || opts.Packages[0] != "./..." {
				t.Fatalf("expected default package ./..., got %#v", opts.Packages)
			}
		})

		t.Run("Custom", func(t *testing.T) {
			opts, code := gotestx.ResolveOptions(
				[]string{"./internal", "./cmd"},
				io.Discard,
				io.Discard,
			)

			if code != gotestx.ExitContinue {
				t.Fatalf("unexpected exit code %d", code)
			}

			if opts == nil {
				t.Fatalf("expected options not nil")
			}

			if len(opts.Packages) != 2 {
				t.Fatalf("expected 2 packages")
			}
		})

	})

	t.Run("Flags", func(t *testing.T) {

		t.Run("Combined", func(t *testing.T) {
			opts, code := gotestx.ResolveOptions(
				[]string{"-coqV"},
				io.Discard,
				io.Discard,
			)

			if code != gotestx.ExitContinue {
				t.Fatalf("unexpected exit code")
			}

			if opts == nil {
				t.Fatalf("expected options not nil")
			}

			if !opts.WithCoverage ||
				!opts.OpenCoverage ||
				!opts.Quiet ||
				!opts.CleanView {
				t.Fatalf("combined flags not parsed correctly: %#v", opts)
			}
		})

		t.Run("Short", func(t *testing.T) {

			t.Run("CleanView", func(t *testing.T) {
				opts, code := gotestx.ResolveOptions(
					[]string{"-V"},
					io.Discard,
					io.Discard,
				)

				if opts == nil {
					t.Fatalf("expected options not nil")
				}

				if code != gotestx.ExitContinue {
					t.Fatalf("expected ExitContinue, got %d", code)
				}

				if !opts.CleanView {
					t.Fatalf("expected CleanView = true")
				}
			})

			t.Run("Coverage", func(t *testing.T) {
				opts, code := gotestx.ResolveOptions(
					[]string{"-c"},
					io.Discard,
					io.Discard,
				)

				if opts == nil {
					t.Fatalf("expected options not nil")
				}

				if code != gotestx.ExitContinue {
					t.Fatalf("expected ExitContinue, got %d", code)
				}

				if !opts.WithCoverage {
					t.Fatalf("expected WithCoverage = true")
				}
			})

			t.Run("Help", func(t *testing.T) {
				var out bytes.Buffer

				opts, code := gotestx.ResolveOptions(
					[]string{"-h"},
					&out,
					io.Discard,
				)

				if opts != nil {
					t.Fatalf("expected nil options")
				}

				if code != gotestx.ExitOK {
					t.Fatalf("expected ExitOK")
				}

				if out.Len() == 0 {
					t.Fatalf("expected help output")
				}
			})

			t.Run("Open", func(t *testing.T) {
				opts, _ := gotestx.ResolveOptions(
					[]string{"-o"},
					io.Discard,
					io.Discard,
				)

				if opts == nil {
					t.Fatalf("expected options not nil")
				}

				if !opts.WithCoverage {
					t.Fatalf("open coverage should imply coverage")
				}
			})

			t.Run("Quiet", func(t *testing.T) {
				opts, code := gotestx.ResolveOptions(
					[]string{"-q"},
					io.Discard,
					io.Discard,
				)

				if opts == nil {
					t.Fatalf("expected options not nil")
				}

				if code != gotestx.ExitContinue {
					t.Fatalf("expected ExitContinue, got %d", code)
				}

				if !opts.Quiet {
					t.Fatalf("expected Quiet = true")
				}
			})

			t.Run("Version", func(t *testing.T) {
				var out bytes.Buffer

				opts, code := gotestx.ResolveOptions(
					[]string{"-v"},
					&out,
					io.Discard,
				)

				if opts != nil {
					t.Fatalf("expected nil options")
				}

				if code != gotestx.ExitOK {
					t.Fatalf("expected ExitOK")
				}

				if out.Len() == 0 {
					t.Fatalf("expected version output")
				}
			})

		})

		t.Run("Long", func(t *testing.T) {
			opts, code := gotestx.ResolveOptions(
				[]string{"--clean-view", "--quiet"},
				io.Discard,
				io.Discard,
			)

			if code != gotestx.ExitContinue {
				t.Fatalf("unexpected exit code")
			}

			if opts == nil {
				t.Fatalf("expected options not nil")
			}

			if !opts.CleanView || !opts.Quiet {
				t.Fatalf("long flags not parsed")
			}
		})

		t.Run("Invalid", func(t *testing.T) {
			var errBuf bytes.Buffer

			opts, code := gotestx.ResolveOptions(
				[]string{"-x"},
				io.Discard,
				&errBuf,
			)

			if opts != nil {
				t.Fatalf("expected nil options")
			}

			if code != gotestx.ExitUsage {
				t.Fatalf("expected ExitUsage")
			}

			if errBuf.Len() == 0 {
				t.Fatalf("expected error output")
			}
		})

	})
}
