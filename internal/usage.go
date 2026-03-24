package gotestx

import (
	"fmt"
	"io"
)

// usage prints the CLI help message describing how to use GoTestX.
func usage(w io.Writer) {
	_, _ = fmt.Fprintf(w, `%s v%s
%s
Author: %s

Usage:
  %s [flags] [packages]

Flags:
  -c, --with-coverage     Enable coverage profile generation (coverage.out)
  -o, --open-coverage     Open coverage report (macOS only, implies -c)
  -q, --quiet             Suppress verbose output (summary only)
  -V, --clean-view        Hide '[no test files]' lines
  -I, --ignore-pattern    Exclude packages matching the given pattern
  -h, --help              Show this help message
  -v, --version           Show version information

Examples:
  %s
  %s -c
  %s -co
  %s -I mock ./...
  %s -q
`,
		ToolName,
		Version,
		Description,
		Author,
		CLIName,
		CLIName,
		CLIName,
		CLIName,
		CLIName,
		CLIName,
	)
}
