package gotestx

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var openCoverageCmd = buildCoverageOpenCommand
var loadIgnore = loadIgnoreFile

var listPackages = func(pkg string) ([]byte, error) {
	return exec.Command("go", "list", pkg).Output()
}

// Run executes GoTestX using the provided CLI arguments.
func Run(args []string, stdout, stderr io.Writer) int {
	opts, exit := ResolveOptions(args, stdout, stderr)
	if exit != ExitContinue {
		return exit
	}

	if opts.OpenCoverage && getGOOS() != "darwin" {
		_, _ = fmt.Fprintln(stderr, "Error: --open-coverage is only supported on macOS.")
		return 1
	}

	inputPkgs := append([]string(nil), opts.Packages...)

	for i, pkg := range inputPkgs {
		if strings.Contains(pkg, "...") {
			continue
		}

		st, err := os.Stat(pkg)
		if err != nil || !st.IsDir() {
			if opts.Quiet {
				_, _ = fmt.Fprintln(stderr, "❌ Tests failed (use without -q to see details)")
			} else {
				_, _ = fmt.Fprintf(stderr, "Error: Package path '%s' does not exist.\n", pkg)
			}

			return 1
		}

		matches, _ := filepath.Glob(filepath.Join(pkg, "*.go"))

		if len(matches) == 0 {
			subMatches, _ := filepath.Glob(filepath.Join(pkg, "**", "*.go"))

			if len(subMatches) > 0 {
				if !opts.Quiet {
					_, _ = fmt.Fprintf(
						stdout,
						"Info: No Go files in '%s', using subpackages instead (%s/...)\n",
						pkg,
						pkg,
					)
				}
				inputPkgs[i] = pkg + "/..."
			} else {
				_, _ = fmt.Fprintf(stderr, "Error: No Go files found in '%s'.\n", pkg)
				return 1
			}
		}
	}

	var expandedPkgs []string

	for _, pkg := range inputPkgs {
		out, err := listPackages(pkg)
		if err != nil {
			_, _ = fmt.Fprintf(stderr, "Error: failed to list packages for '%s': %v\n", pkg, err)
			return 1
		}

		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		expandedPkgs = append(expandedPkgs, lines...)
	}

	filePatterns, err := loadIgnore()
	if err != nil {
		_, _ = fmt.Fprintf(stderr, "Error loading .gotestxignore: %v\n", err)
		return 1
	}

	patterns := append(filePatterns, opts.Ignore...)

	packages := expandedPkgs

	if len(patterns) > 0 {
		filtered := filterPackages(packages, patterns)

		if len(filtered) == 0 {
			if !opts.Quiet {
				_, _ = fmt.Fprintln(stdout, "No packages to test after applying ignore rules.")
			}
			return ExitOK
		}

		packages = filtered
	}

	var cmd Command

	if opts.WithCoverage {
		if !opts.Quiet {
			_, _ = fmt.Fprintln(stdout, "Running tests with coverage...")
		}

		args := buildCoverageArgs(packages)
		cmd = commandRunner("go", args...)

	} else {
		if !opts.Quiet {
			_, _ = fmt.Fprintln(stdout, "Running tests...")
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

	err = cmd.Run()

	if opts.CleanView && !opts.Quiet {
		FilterCleanViewOutput(&buf, stdout)
	}

	if opts.Quiet {
		return HandleQuietOutput(buf.String(), err, opts, stdout, stderr)
	}

	if err != nil {
		_, _ = fmt.Fprintf(stderr, "Error: go test failed: %v\n", err)
		return 1
	}

	if opts.WithCoverage && !opts.Quiet {
		_, _ = fmt.Fprintln(stdout, "Coverage: coverage.out (use -o to open)")
		_, _ = fmt.Fprintln(stdout, "Run 'go tool cover -html=coverage.out' to view it")
	}

	if opts.WithCoverage && opts.OpenCoverage {
		if !opts.Quiet {
			_, _ = fmt.Fprintln(stdout, "Opening coverage report in browser...")
		}

		openCmd := openCoverageCmd()
		openCmd.SetStdout(stdout)
		openCmd.SetStderr(stderr)

		if err := openCmd.Run(); err != nil {
			_, _ = fmt.Fprintf(stderr, "Error: failed to open coverage report: %v\n", err)
			return 1
		}
	}

	return ExitOK
}
