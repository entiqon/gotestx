package gotestx

import (
	"fmt"
	"io"
	"strings"
)

// HandleQuietOutput processes buffered test output when quiet mode is enabled.
func HandleQuietOutput(
	buf string,
	err error,
	opts *Options,
	stdout io.Writer,
	stderr io.Writer,
) int {
	lines := strings.Split(strings.TrimSpace(buf), "\n")

	if err != nil {
		fmt.Fprintln(stderr, "❌ Tests failed (use without -q to see details)")
		return 1
	}

	if opts.WithCoverage {
		for i := len(lines) - 1; i >= 0; i-- {
			if strings.Contains(lines[i], "coverage:") {
				fmt.Fprintln(stdout, lines[i])
				break
			}
		}
	} else {
		fmt.Fprintln(stdout, "✅ Tests finished successfully")
	}

	return ExitOK
}
