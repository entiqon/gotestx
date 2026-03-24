# GoTestX

> A developer-friendly test runner for Go with coverage, filtering, and clean output.

[![Official Repository](https://img.shields.io/badge/Repository-Entiqon%20Labs-blue?logo=github)](https://github.com/entiqon/gotestx)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/entiqon/gotestx)](https://github.com/entiqon/gotestx/releases)
[![Build Status](https://github.com/entiqon/gotestx/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/gotestx/actions)
[![Codecov](https://codecov.io/gh/entiqon/gotestx/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/gotestx)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/gotestx)](https://goreportcard.com/report/github.com/entiqon/gotestx)
[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/gotestx.svg)](https://pkg.go.dev/github.com/entiqon/gotestx)
[![License](https://img.shields.io/github/license/entiqon/gotestx)](LICENSE)

**Official repository (canonical source):** https://github.com/entiqon/gotestx

> [!CAUTION]
> **OFFICIAL INSTALLATION ONLY**
> Please ensure you are using the official **Entiqon/GoTestX** repository.
> Never download binaries from untrusted forks. Use only the official releases.

---

## Usage

```bash
gotestx [flags] [packages]
```

If no packages are provided, GoTestX defaults to:

    ./...

---

## Flags

| Flag                     | Description                                         |
|--------------------------|-----------------------------------------------------|
| `-c`, `--with-coverage`  | Enable coverage profile generation (`coverage.out`) |
| `-o`, `--open-coverage`  | Open coverage report (macOS only, implies `-c`)     |
| `-q`, `--quiet`          | Suppress verbose output (summary only)              |
| `-V`, `--clean-view`     | Hide `[no test files]` lines                        |
| `-I`, `--ignore-pattern` | Exclude packages using tree-based patterns          |
| `-h`, `--help`           | Show help                                           |
| `-v`, `--version`        | Show version                                        |

---

## Execution Behavior

### Package Resolution

- Packages are resolved using `go list`
- If no packages are provided → `./...` is used
- Invalid paths fail fast

---

### Package Filtering

Packages can be excluded using:

- CLI flags: `-I <pattern>`
- `.gotestxignore` file

#### Pattern Rules

Patterns follow **tree-based matching**:

| Pattern           | Meaning                          |
|-------------------|----------------------------------|
| `mock`            | matches any segment named `mock` |
| `outport/testkit` | matches exact subpath            |

- Matching is **segment-based**
- Not substring-based
- Not glob-based

---

### Execution Flow

```
go list ./...
      ↓
load .gotestxignore
      ↓
merge CLI patterns
      ↓
filter packages
      ↓
go test
```

---

### Filtering Behavior

- Applied **after package discovery**
- Matching packages are removed before execution
- If all packages are excluded:

```bash
No packages to test after applying ignore rules.
```

---

## Output Modes

### Default

```bash
gotestx
```

```
Running tests...
ok   pkg/a
ok   pkg/b
```

---

### Coverage

```bash
gotestx -c
```

```
Running tests with coverage...
ok   pkg/a
ok   pkg/b
Coverage: coverage.out
```

---

### Quiet Mode

```bash
gotestx -q
```

```
Tests finished successfully
```

---

### Clean View

```bash
gotestx -V
```

Removes lines like:

```
? pkg [no test files]
```

---

## Examples

```bash
gotestx
gotestx -c
gotestx -cq
gotestx -I mock ./...
gotestx -I mock -I testkit ./...
```

---

## Notes

- `--open-coverage` is supported only on macOS
- Flags can be combined (e.g. `-cq`, `-co`)
