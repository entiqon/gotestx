// Package gotestx provides the internal runtime used by GoTestX.
//
// The package implements the core execution pipeline responsible for:
//
//   - parsing CLI options
//   - orchestrating `go test` execution
//   - handling coverage generation
//   - filtering and formatting test output
//   - excluding packages via ignore rules
//
// The runtime is intentionally placed under Go's `internal` directory to
// prevent external modules from importing unstable implementation details.
// Only the CLI entrypoint (main.go) should interact with this package.
//
// The Run function acts as the orchestrator and delegates specific tasks to
// smaller focused components.
//
// # Components
//
//	options.go     — CLI argument parsing
//	run.go         — execution orchestrator
//	coverage.go    — coverage helpers
//	quiet.go       — quiet-mode output handling
//	cleanview.go   — output filtering
//	ignore.go      — package exclusion and filtering logic
//	ignorefile.go  — .gotestxignore loading
//	runtime.go     — runtime helpers
//	command.go     — command abstraction
//	usage.go       — help output
//	version.go     — version information
//	exitcodes.go   — CLI exit semantics
//
// # Exclusion
//
// GoTestX supports excluding packages from execution using:
//
//   - CLI flag: -I <pattern>
//   - Ignore file: .gotestxignore
//
// Patterns follow a tree-based syntax:
//
//	mock              → excludes any path segment named "mock"
//	outport/testkit   → excludes matching subpaths
//
// Filtering is applied after package discovery and before test execution.
// If all packages are excluded, execution exits successfully with a message.
//
// # Stability
//
// This package contains internal implementation details and may change
// without notice between versions of GoTestX.
package gotestx
