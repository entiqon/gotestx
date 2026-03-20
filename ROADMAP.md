# GoTestX Roadmap

GoTestX is a lightweight CLI wrapper around `go test` that improves
developer productivity by providing better output control, coverage
utilities, and test execution helpers.

This roadmap outlines planned improvements while maintaining GoTestX's
core philosophy:

-   Minimal abstraction over `go test`
-   Fully testable CLI
-   Zero external dependencies
-   Developer-focused ergonomics

---

## v1.1.0

* Coverage execution
* Coverage report generation
* Coverage HTML viewer (macOS)
* Quiet output mode
* Clean output mode (hide `[no test files]`)
* Package validation
* Default recursive package execution (`./...`)
* Fully testable runtime (`Run(args, stdout, stderr)`)

```
gotestx -c
gotestx -q
gotestx -V
gotestx ./internal/...
```

## v1.1.3 (Current version)

Core capabilities:

-   Coverage execution
-   Coverage report generation
-   Coverage HTML viewer (macOS)
-   Quiet output mode
-   Clean output mode (hide `[no test files]`)
-   Package validation
-   Default recursive package execution (`./...`)
-   Fully testable runtime (`Run(args, stdout, stderr)`)

Example:

    gotestx -c
    gotestx -q
    gotestx -V
    gotestx ./internal/...

---

# Short-Term Roadmap

## v1.2.0 --- Package Ignore Support

Introduce package exclusion during test execution using CLI flags and
`.gotestxignore`.

### Flags

    -I, --ignore-pattern

Exclude packages from execution using tree-based patterns.

### Examples

    gotestx -I mock
    gotestx -I mock -I testkit
    gotestx ./... -I mock

### Behavior

- Patterns are applied after package discovery (`go list`)
- Matching is segment-based (tree syntax), not glob-based
- CLI patterns are merged with `.gotestxignore`
- Packages matching any pattern are removed before execution
- If all packages are excluded:

    No packages to test after applying ignore rules.

### Pattern Rules

Patterns are evaluated against path segments:

  Pattern              Meaning
  -------------------- ------------------------------------
  `mock`               Matches any segment named "mock"
  `outport/testkit`    Matches exact subpath

### `.gotestxignore`

Ignore rules can be defined in a `.gotestxignore` file at the project root.

Example:

    # Ignore integration tests
    apitest

    # Ignore generated code
    internal/generated

    # Ignore helper packages
    testkit

### Execution Flow

    go list ./...
          ↓
    load .gotestxignore
          ↓
    merge CLI patterns
          ↓
    filter packages
          ↓
    go test

---

## v1.3.0 --- Coverage Insights

Improve visibility of coverage results.

### New Flags

    -s, --summary

Displays summarized coverage results extracted from `coverage.out`.

Example:

    gotestx -c -s

Output:

    Coverage Summary
    ----------------
    total: (statements) 91.2%

---

## v1.4.0 --- Execution Controls

### New Flags

    -p, --parallel N

Maps to:

    go test -parallel=N

---

## v1.5.0 --- Fail Fast Mode

### New Flags

    -f, --fail-fast

Maps to:

    go test -failfast

---

## v1.6.0 --- Benchmark Support

### New Flags

    -b, --bench

Maps to:

    go test -bench=.

---

# Project

GoTestX is part of the Entiqon developer tooling ecosystem.
