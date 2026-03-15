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

Usage: %s [options] [packages]

Options:
  -c, --with-coverage   Run tests with coverage report generation (coverage.out)
  -o, --open-coverage   Open coverage report in browser (macOS only, implies -c)
  -q, --quiet           Suppress verbose test chatter (only summary shown)
  -V, --clean-view      Suppress 'no test files' lines for cleaner output
  -h, --help            Show this help
  -v, --version         Show version info
`,
		ToolName,
		Version,
		Description,
		Author,
		ToolName,
	)
}
