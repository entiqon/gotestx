// Package gotestx provides the internal runtime used by GoTestX.
//
// The package implements the core execution pipeline responsible for:
//
//   - parsing CLI options
//   - orchestrating `go test` execution
//   - handling coverage generation
//   - filtering and formatting test output
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
//	runtime.go     — runtime helpers
//	command.go     — command abstraction
//	usage.go       — help output
//	version.go     — version information
//	exitcodes.go   — CLI exit semantics
//
// # Stability
//
// This package contains internal implementation details and may change
// without notice between versions of GoTestX.
package gotestx
