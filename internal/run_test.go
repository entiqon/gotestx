package gotestx

import (
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func withDeps(t *testing.T, fn func()) {
	oldList := listPackages
	oldCmd := commandRunner
	oldOpen := openCoverageCmd
	oldOS := getGOOS
	oldIgnore := loadIgnore

	t.Cleanup(func() {
		listPackages = oldList
		commandRunner = oldCmd
		openCoverageCmd = oldOpen
		getGOOS = oldOS
		loadIgnore = oldIgnore
	})

	fn()
}

type fakeCommand struct {
	runErr error
	output string
	stdout io.Writer
	stderr io.Writer
}

func (f *fakeCommand) Run() error {
	if f.stdout != nil {
		_, _ = f.stdout.Write([]byte(f.output))
	}
	return f.runErr
}

func (f *fakeCommand) SetStdout(w io.Writer) { f.stdout = w }
func (f *fakeCommand) SetStderr(w io.Writer) { f.stderr = w }

func run(t *testing.T, args ...string) (string, string, int) {
	t.Helper()

	var out, err bytes.Buffer
	code := Run(args, &out, &err)

	return out.String(), err.String(), code
}

func TestRun(t *testing.T) {
	t.Run("Packages", func(t *testing.T) {
		t.Run("Invalid", func(t *testing.T) {
			_, err, code := run(t, "./does-not-exist")

			if code == ExitOK {
				t.Fatalf("expected failure")
			}
			if !strings.Contains(err, "does not exist") {
				t.Fatalf("expected error message")
			}
		})

		t.Run("Success", func(t *testing.T) {
			withDeps(t, func() {
				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				_, _, code := run(t, "./...")

				if code != ExitOK {
					t.Fatalf("expected success")
				}
			})
		})
	})

	t.Run("NoGoFiles", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			tmp := t.TempDir()

			_, err, code := run(t, tmp)

			if code == 0 {
				t.Fatalf("expected failure")
			}
			if !strings.Contains(err, "No Go files") {
				t.Fatalf("expected error")
			}
		})

		t.Run("ButSubpackages", func(t *testing.T) {
			withDeps(t, func() {
				tmp := t.TempDir()

				root := filepath.Join(tmp, "pkg")
				sub := filepath.Join(root, "sub")

				_ = os.MkdirAll(sub, 0o755)
				_ = os.WriteFile(filepath.Join(sub, "main.go"), []byte("package sub"), 0o644)

				listPackages = func(string) ([]byte, error) {
					return []byte("pkg/sub"), nil
				}

				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				out, _, code := run(t, root)

				if code != ExitOK {
					t.Fatalf("expected success")
				}
				if !strings.Contains(out, "using subpackages") {
					t.Fatalf("expected fallback message")
				}
			})
		})

		t.Run("Anywhere", func(t *testing.T) {
			tmp := t.TempDir()

			root := filepath.Join(tmp, "empty")
			_ = os.MkdirAll(root, 0o755)

			_, err, code := run(t, root)

			if code == 0 {
				t.Fatalf("expected failure")
			}
			if !strings.Contains(err, "No Go files found") {
				t.Fatalf("expected error message")
			}
		})

		t.Run("AnywhereQuiet", func(t *testing.T) {
			tmp := t.TempDir()

			root := filepath.Join(tmp, "empty")
			_ = os.MkdirAll(root, 0o755)

			_, _, code := run(t, "-q", root)

			if code == 0 {
				t.Fatalf("expected failure")
			}
		})
	})

	t.Run("CleanView", func(t *testing.T) {
		withDeps(t, func() {
			commandRunner = func(string, ...string) Command {
				return &fakeCommand{output: "ok"}
			}

			_, _, code := run(t, "-V")

			if code != ExitOK {
				t.Fatalf("expected success")
			}
		})
	})

	t.Run("Coverage", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			withDeps(t, func() {
				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				_, _, code := run(t, "-c")

				if code != ExitOK {
					t.Fatalf("expected success")
				}
			})
		})

		t.Run("OpenCoverage", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				withDeps(t, func() {
					getGOOS = func() string { return "darwin" }

					listPackages = func(string) ([]byte, error) {
						return []byte("pkg/a"), nil
					}

					commandRunner = func(string, ...string) Command {
						return &fakeCommand{}
					}

					openCoverageCmd = func() Command {
						return &fakeCommand{}
					}

					out, _, code := run(t, "-c", "-o")

					if code != ExitOK {
						t.Fatalf("expected success")
					}
					if !strings.Contains(out, "Opening coverage report") {
						t.Fatalf("expected message")
					}
				})
			})

			t.Run("NonDarwin", func(t *testing.T) {
				withDeps(t, func() {
					getGOOS = func() string { return "linux" }

					_, err, code := run(t, "-o")

					if code == 0 {
						t.Fatalf("expected failure")
					}
					if !strings.Contains(err, "only supported on macOS") {
						t.Fatalf("expected macOS error")
					}
				})
			})

			t.Run("Failure", func(t *testing.T) {
				withDeps(t, func() {
					getGOOS = func() string { return "darwin" }

					listPackages = func(string) ([]byte, error) {
						return []byte("pkg/a"), nil
					}

					commandRunner = func(string, ...string) Command {
						return &fakeCommand{}
					}

					openCoverageCmd = func() Command {
						return &fakeCommand{runErr: errors.New("fail")}
					}

					_, err, code := run(t, "-c", "-o")

					if code == 0 {
						t.Fatalf("expected failure")
					}
					if !strings.Contains(err, "failed to open coverage report") {
						t.Fatalf("expected error")
					}
				})
			})
		})
	})

	t.Run("Help", func(t *testing.T) {
		out, _, code := run(t, "-h")

		if code != 0 {
			t.Fatalf("expected exit 0")
		}
		if !strings.Contains(out, "Usage") {
			t.Fatalf("expected usage")
		}
	})

	t.Run("IgnoreFile", func(t *testing.T) {
		withDeps(t, func() {
			loadIgnore = func() ([]string, error) {
				return nil, errors.New("boom")
			}

			_, err, code := run(t, "./...")

			if code == 0 {
				t.Fatalf("expected failure")
			}
			if !strings.Contains(err, "Error loading .gotestxignore") {
				t.Fatalf("expected ignore error")
			}
		})
	})

	t.Run("Quiet", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			withDeps(t, func() {
				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				_, _, code := run(t, "-q")

				if code != ExitOK {
					t.Fatalf("expected success")
				}
			})
		})

		t.Run("InvalidPackage", func(t *testing.T) {
			_, err, code := run(t, "-q", "./does-not-exist")

			if code == 0 {
				t.Fatalf("expected failure")
			}
			if !strings.Contains(err, "Tests failed") {
				t.Fatalf("expected quiet error")
			}
			if strings.Contains(err, "does not exist") {
				t.Fatalf("should not leak details")
			}
		})
	})

	t.Run("Version", func(t *testing.T) {
		out, _, code := run(t, "-v")

		if code != 0 {
			t.Fatalf("expected exit 0")
		}
		if !strings.Contains(out, Version) {
			t.Fatalf("expected version")
		}
	})

	t.Run("Filter", func(t *testing.T) {
		t.Run("All", func(t *testing.T) {
			withDeps(t, func() {
				listPackages = func(string) ([]byte, error) {
					return []byte("pkg/mock\npkg/mock"), nil
				}

				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				out, _, code := run(t, "./...", "-I", "mock")

				if code != ExitOK {
					t.Fatalf("expected success")
				}
				if !strings.Contains(out, "No packages") {
					t.Fatalf("expected message")
				}
			})
		})

		t.Run("Partial", func(t *testing.T) {
			withDeps(t, func() {
				listPackages = func(string) ([]byte, error) {
					return []byte("pkg/a\npkg/mock\npkg/b"), nil
				}

				commandRunner = func(string, ...string) Command {
					return &fakeCommand{}
				}

				out, _, code := run(t, "./...", "-I", "mock")

				if code != ExitOK {
					t.Fatalf("expected success")
				}
				if !strings.Contains(out, "Running tests") {
					t.Fatalf("expected filtered packages")
				}
			})
		})
	})

	t.Run("GoListFailure", func(t *testing.T) {
		t.Run("Default", func(t *testing.T) {
			withDeps(t, func() {
				listPackages = func(string) ([]byte, error) {
					return nil, errors.New("boom")
				}

				_, err, code := run(t, "./...")

				if code == 0 {
					t.Fatalf("expected failure")
				}
				if !strings.Contains(err, "failed to list packages") {
					t.Fatalf("expected go list error")
				}
			})
		})

		t.Run("NonQuiet", func(t *testing.T) {
			withDeps(t, func() {
				listPackages = func(string) ([]byte, error) {
					return []byte("pkg/a"), nil
				}

				commandRunner = func(string, ...string) Command {
					return &fakeCommand{runErr: errors.New("fail")}
				}

				_, err, code := run(t, "./...")

				if code == 0 {
					t.Fatalf("expected failure")
				}
				if !strings.Contains(err, "go test failed") {
					t.Fatalf("expected go test failure")
				}
			})
		})
	})
}
