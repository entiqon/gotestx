package gotestx

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadIgnoreFile(t *testing.T) {
	t.Run("FileDoesNotExist", func(t *testing.T) {
		tmp := t.TempDir()

		old, _ := os.Getwd()
		defer func(dir string) {
			_ = os.Chdir(dir)
		}(old)
		_ = os.Chdir(tmp)

		patterns, err := loadIgnoreFile()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if patterns != nil {
			t.Fatalf("expected nil, got %v", patterns)
		}
	})

	t.Run("ValidFile", func(t *testing.T) {
		tmp := t.TempDir()

		content := `
# comment
mock

outport/testkit

   # another comment
service
`

		write(t, tmp, ".gotestxignore", content)

		old, _ := os.Getwd()
		defer func(dir string) {
			_ = os.Chdir(dir)
		}(old)
		_ = os.Chdir(tmp)

		patterns, err := loadIgnoreFile()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expect := []string{"mock", "outport/testkit", "service"}

		if len(patterns) != len(expect) {
			t.Fatalf("expected %v, got %v", expect, patterns)
		}

		for i := range patterns {
			if patterns[i] != expect[i] {
				t.Fatalf("expected %v, got %v", expect, patterns)
			}
		}
	})
}

func TestLoadIgnoreFile_OpenError(t *testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	defer func(dir string) {
		_ = os.Chdir(dir)
	}(oldWd)

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	// THIS is the missing part
	if err := os.WriteFile(".gotestxignore", []byte("mock"), 0o644); err != nil {
		t.Fatal(err)
	}

	old := openFile
	defer func() { openFile = old }()

	openFile = func(string) (*os.File, error) {
		return nil, os.ErrPermission
	}

	patterns, err := loadIgnoreFile()
	if err == nil {
		t.Fatal("expected error")
	}
	if patterns != nil {
		t.Fatalf("expected nil patterns, got %v", patterns)
	}
}

type badScanner struct{}

func (badScanner) Scan() bool   { return false }
func (badScanner) Text() string { return "" }
func (badScanner) Err() error   { return io.ErrUnexpectedEOF }

func TestLoadIgnoreFile_ScannerError(t *testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	defer func(dir string) {
		_ = os.Chdir(dir)
	}(oldWd)
	_ = os.Chdir(tmp)

	if err := os.WriteFile(".gotestxignore", []byte("mock"), 0o644); err != nil {
		t.Fatal(err)
	}

	oldScanner := newScanner
	defer func() { newScanner = oldScanner }()

	newScanner = func(io.Reader) scanner {
		return badScanner{}
	}

	_, err := loadIgnoreFile()
	if err == nil {
		t.Fatal("expected scanner error")
	}
}

func write(t *testing.T, root, path, content string) {
	t.Helper()

	full := filepath.Join(root, path)

	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
