package gotestx

import (
	"bufio"
	"io"
	"os"
	"strings"
)

var openFile = os.Open

type scanner interface {
	Scan() bool
	Text() string
	Err() error
}

var newScanner = func(r io.Reader) scanner {
	return bufio.NewScanner(r)
}

// loadIgnoreFile reads ignore patterns from a `.gotestxignore` file
// in the current working directory.
//
// Each non-empty, non-comment line is treated as a pattern.
//
// Patterns follow tree syntax:
//   - "mock" matches any segment named "mock"
//   - "outport/testkit" matches exact contiguous path
//
// No globbing or wildcard expansion is supported.
func loadIgnoreFile() ([]string, error) {
	const file = ".gotestxignore"

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, nil
	}

	f, err := openFile(file)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	var patterns []string

	scanner := newScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		patterns = append(patterns, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return patterns, nil
}
