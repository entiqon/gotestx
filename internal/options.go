package gotestx

import (
	"fmt"
	"io"
	"strings"
)

// Options represents the resolved CLI configuration used
// to execute GoTestX.
type Options struct {
	// WithCoverage enables coverage profile generation.
	WithCoverage bool

	// OpenCoverage opens the coverage report after tests complete.
	OpenCoverage bool

	// Quiet suppresses verbose test output.
	Quiet bool

	// CleanView hides lines reporting "[no test files]".
	CleanView bool

	// Packages contains the packages that should be tested.
	Packages []string
}

// ResolveOptions parses CLI arguments and resolves them into an
// Options configuration.
//
// It performs the following:
//
//   - Parses short and long flags
//   - Supports combined short flags (e.g. -cqV)
//   - Resolves positional arguments as Go packages
//   - Applies defaults when no packages are provided
//   - Normalizes dependent flags (e.g. --open-coverage implies -c)
//
// Return values:
//
//	*Options — resolved configuration
//	int      — exit code indicator
//
// Exit semantics:
//
//	ExitContinue → normal execution should continue
//	ExitOK       → command completed (help/version printed)
//	ExitUsage    → CLI usage error
func ResolveOptions(args []string, stdout, stderr io.Writer) (*Options, int) {
	opts := &Options{}

	for i := 0; i < len(args); i++ {
		arg := args[i]

		switch {
		case arg == "-h" || arg == "--help":
			usage(stdout)
			return nil, ExitOK

		case arg == "-v" || arg == "--version":
			versionInfo(stdout)
			return nil, ExitOK

		case arg == "-c" || arg == "--with-coverage":
			opts.WithCoverage = true

		case arg == "-o" || arg == "--open-coverage":
			opts.OpenCoverage = true

		case arg == "-q" || arg == "--quiet":
			opts.Quiet = true

		case arg == "-V" || arg == "--clean-view":
			opts.CleanView = true

		case strings.HasPrefix(arg, "-"):
			flags := arg[1:]
			for _, f := range flags {
				switch f {
				case 'c':
					opts.WithCoverage = true
				case 'o':
					opts.OpenCoverage = true
				case 'q':
					opts.Quiet = true
				case 'V':
					opts.CleanView = true
				default:
					_, _ = fmt.Fprintf(stderr, "Error: Unknown short option: -%c\n", f)
					usage(stderr)
					return nil, ExitUsage
				}
			}

		default:
			opts.Packages = append(opts.Packages, arg)
		}
	}

	if len(opts.Packages) == 0 {
		opts.Packages = []string{"./..."}
	}

	if opts.OpenCoverage && !opts.WithCoverage {
		opts.WithCoverage = true
	}

	return opts, ExitContinue
}
