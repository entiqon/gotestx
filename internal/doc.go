// Package gotestx provides the core runtime implementation used by GoTestX.
//
// This package contains the internal components responsible for parsing CLI
// options, orchestrating test execution, handling coverage generation, and
// managing CLI output behavior.
//
// The package is intentionally placed under Go's `internal` directory to
// prevent external modules from importing unstable implementation details.
// Only the CLI entrypoint (main.go) should interact with this package.
//
// # Responsibilities
//
// The internal runtime is responsible for:
//
//   - Parsing CLI flags and arguments
//   - Resolving execution options
//   - Running `go test` commands
//   - Handling coverage generation and reporting
//   - Filtering and formatting command output
//   - Providing CLI usage and version information
//
// # Execution Flow
//
// The typical execution flow of GoTestX is:
//
//	main.go
//	   ↓
//	ResolveOptions
//	   ↓
//	command execution
//	   ↓
//	run tests
//	   ↓
//	process coverage and output
//
// # Exit Semantics
//
// CLI argument resolution returns one of three exit codes:
//
//	ExitContinue  — normal execution should continue
//	ExitOK        — command completed (help/version printed)
//	ExitUsage     — invalid CLI usage
//
// These codes allow the CLI entrypoint to decide whether to continue the
// execution pipeline or terminate early.
//
// # Stability
//
// The APIs in this package are considered internal implementation details and
// may change without notice between versions of GoTestX.
package gotestx
