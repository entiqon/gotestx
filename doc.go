// Package main provides the CLI entrypoint for GoTestX.
//
// GoTestX is a developer-friendly wrapper around `go test` that enhances
// the testing experience with:
//
//   - coverage generation and reporting
//   - package exclusion via ignore rules
//   - quiet and clean output modes
//   - consistent CLI behavior across projects
//
// # Features
//
// GoTestX extends the standard Go testing workflow with:
//
//   - Coverage: generate and optionally open HTML coverage reports
//   - Exclusion: skip packages using CLI flags or `.gotestxignore`
//   - Quiet mode: concise output for CI and scripting
//   - Clean view: remove noisy `[no test files]` lines
//
// # Usage
//
//	Run tests across the module:
//
//	    gotestx
//
//	Run with coverage:
//
//	    gotestx -c
//
//	Open coverage report (macOS only):
//
//	    gotestx -co
//
//	Exclude packages:
//
//	    gotestx -I mock -I testkit ./...
//
//	Quiet mode:
//
//	    gotestx -q
//
// # Exclusion
//
// Packages can be excluded using:
//
//   - CLI flag: `-I <pattern>`
//   - Ignore file: `.gotestxignore`
//
// Patterns use a tree-based syntax:
//
//	mock              → excludes any path segment named "mock"
//	outport/testkit   → excludes matching subpaths
//
// If all packages are excluded, GoTestX exits successfully with:
//
//	No packages to test after applying ignore rules.
//
// # Notes
//
// - Opening coverage reports is supported only on macOS.
// - If no packages are provided, GoTestX defaults to `./...`.
//
// This package is intended to be used as a CLI tool via `main.go`.
package main
