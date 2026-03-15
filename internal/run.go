package gotestx

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Run executes GoTestX using the provided CLI arguments.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, exit := ResolveOptions(args, stdout, stderr)
	if exit != ExitContinue {
		return exit
	}

	packages := append([]string(nil), opts.Packages...)

	for i, pkg := range packages {
		if strings.Contains(pkg, "...") {
			continue
		}

		st, err := os.Stat(pkg)
		if err != nil || !st.IsDir() {
			if opts.Quiet {
				fmt.Fprintln(stderr, "❌ Tests failed (use without -q to see details)")
			} else {
				fmt.Fprintf(stderr, "Error: Package path '%s' does not exist.\n", pkg)
			}

			return 1
		}

		matches, _ := filepath.Glob(filepath.Join(pkg, "*.go"))

		if len(matches) == 0 {
			subMatches, _ := filepath.Glob(filepath.Join(pkg, "**", "*.go"))

			if len(subMatches) > 0 {
				if !opts.Quiet {
					fmt.Fprintf(
						stdout,
						"Info: No Go files in '%s', using subpackages instead (%s/...)\n",
						pkg,
						pkg,
					)
				}

				packages[i] = pkg + "/..."
			} else {

				fmt.Fprintf(stderr, "Error: No Go files found in '%s'.\n", pkg)
				return 1
			}
		}
	}

	if opts.OpenCoverage && getGOOS() != "darwin" {
		fmt.Fprintln(stderr, "Error: --open-coverage is only supported on macOS.")
		return 1
	}

	pkgList := strings.Join(packages, " ")

	var cmd Command

	if opts.WithCoverage {
		if !opts.Quiet {
			fmt.Fprintf(stdout, "Running tests with coverage across: %s\n", pkgList)
		}

		args := BuildCoverageArgs(packages)

		cmd = commandRunner("go", args...)

	} else {
		if !opts.Quiet {
			fmt.Fprintf(stdout, "Running tests normally across: %s\n", pkgList)
		}

		args := append([]string{"test"}, packages...)

		cmd = commandRunner("go", args...)
	}

	var buf bytes.Buffer

	captureOutput := opts.CleanView || opts.Quiet

	if captureOutput {

		cmd.SetStdout(&buf)
		cmd.SetStderr(&buf)

	} else {

		cmd.SetStdout(stdout)
		cmd.SetStderr(stderr)
	}

	err := cmd.Run()

	if opts.CleanView && !opts.Quiet {
		FilterCleanViewOutput(&buf, stdout)
	}

	if opts.Quiet {
		return HandleQuietOutput(buf.String(), err, opts, stdout, stderr)
	}

	if err != nil && !opts.Quiet {
		fmt.Fprintf(stderr, "Error: go test failed: %v\n", err)
		return 1
	}

	if opts.WithCoverage && !opts.Quiet {

		fmt.Fprintln(stdout, "Coverage report saved as coverage.out")
		fmt.Fprintln(stdout, "Run 'go tool cover -html=coverage.out' to view it")
	}

	if opts.WithCoverage && opts.OpenCoverage {
		if !opts.Quiet {
			fmt.Fprintln(stdout, "Opening coverage report in browser...")
		}

		openCmd := BuildCoverageOpenCommand()

		openCmd.SetStdout(stdout)
		openCmd.SetStderr(stderr)

		if err := openCmd.Run(); err != nil {

			fmt.Fprintf(stderr, "Error: failed to open coverage report: %v\n", err)
			return 1
		}
	}

	return ExitOK
}
