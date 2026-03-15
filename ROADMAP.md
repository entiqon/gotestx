# GoTestX Roadmap

GoTestX is a lightweight CLI wrapper around `go test` that improves developer productivity by providing better output control, coverage utilities, and test execution helpers.

This roadmap outlines planned improvements while maintaining GoTestX’s core philosophy:

* Minimal abstraction over `go test`
* Fully testable CLI
* Zero external dependencies
* Developer-focused ergonomics

---

# Current Version

## v1.1.0 (Current)

Core capabilities:

* Coverage execution
* Coverage report generation
* Coverage HTML viewer (macOS)
* Quiet output mode
* Clean output mode (hide `[no test files]`)
* Package validation
* Default recursive package execution (`./...`)
* Fully testable runtime (`Run(args, stdout, stderr)`)

Example:

```
gotestx -c
gotestx -q
gotestx -V
gotestx ./internal/...
```

---

# Short-Term Roadmap

## v1.2.0 — Coverage Insights

Improve visibility of coverage results.

### New Flags

```
-s, --summary
```

Displays summarized coverage results extracted from `coverage.out`.

Example:

```
gotestx -c -s
```

Output:

```
Coverage Summary
----------------
total: (statements) 91.2%
```

Future enhancement may include package-level coverage aggregation.

---

## v1.3.0 — Execution Controls

Add control over test execution behavior.

### New Flags

```
-p, --parallel N
```

Maps to:

```
go test -parallel=N
```

Example:

```
gotestx -p 8
```

---

## v1.4.0 — Fail Fast Mode

Allow early test failure detection.

### New Flags

```
-f, --fail-fast
```

Maps to:

```
go test -failfast
```

Example:

```
gotestx -f
```

---

## v1.5.0 — Benchmark Support

Enable benchmark execution from the CLI.

### New Flags

```
-b, --bench
```

Maps to:

```
go test -bench=.
```

Example:

```
gotestx -b
```

---

# Mid-Term Roadmap

## v1.6.0 — JSON Output Mode

Support machine-readable output for CI systems.

### New Flags

```
--json
```

Maps to:

```
go test -json
```

Example:

```
gotestx --json
```

This allows GoTestX to integrate with CI dashboards and external tooling.

---

## v1.7.0 — Coverage Ranking

Highlight test coverage weaknesses across packages.

### New Flags

```
--rank
```

Example:

```
gotestx -c --rank
```

Output:

```
Coverage Ranking
----------------
1. pkg/utils       100%
2. pkg/api         92.1%
3. pkg/service     88.7%
4. pkg/repository  71.2%
```

---

# Long-Term Vision

The long-term goal is to evolve GoTestX into a **developer productivity tool for Go testing**, while remaining simple and dependency-free.

Potential features:

### Coverage Improvements

* Package coverage summaries
* Coverage trend tracking
* Coverage diff between commits

### Developer Insights

* Slowest test detection
* Flaky test detection
* Test runtime statistics

### CI/CD Support

* Structured output formats
* GitHub Actions integration helpers
* Coverage threshold enforcement

Example:

```
gotestx --min-coverage 85
```

---

# Architectural Improvements

To support future growth, the codebase may be modularized.

Current structure:

```
gotestx.go
```

Proposed structure:

```
gotestx/
    run.go
    flags.go
    coverage.go
    summary.go
    usage.go
    version.go
```

This refactor would improve maintainability without changing public behavior.

---

# Contribution Guidelines (Planned)

Future documentation will include:

* Contribution guide
* Test strategy documentation
* CLI design principles
* Release process

---

# Release Philosophy

GoTestX follows a **feature-incremental release strategy**:

* Small improvements per release
* No breaking CLI changes without major version bump
* Maintain compatibility with `go test`

---

# Project

GoTestX is part of the **Entiqon developer tooling ecosystem**.
