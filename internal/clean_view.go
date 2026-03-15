package gotestx

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// FilterCleanViewOutput removes noisy lines from Go test output when
// CleanView mode is enabled. Currently, it filters lines containing
// "[no test files]" while preserving all other output.
func FilterCleanViewOutput(buf io.Reader, stdout io.Writer) {
	scanner := bufio.NewScanner(buf)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "[no test files]") {
			continue
		}

		fmt.Fprintln(stdout, line)
	}
}
