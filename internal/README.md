# internal

The `internal` package contains the core runtime implementation used by
**GoTestX**.

This package is not intended to be imported by external modules. It
provides the internal orchestration logic that powers the CLI command
exposed in `main.go`.

Responsibilities of this package include:

- CLI option parsing
- command execution
- runtime helpers
- coverage handling
- output filtering
- package exclusion via ignore rules
- quiet execution summaries
- exit semantics

---

## CLI Exit Semantics

`ResolveOptions` returns an exit code that controls how the CLI should
behave after argument parsing.

| Exit Code      | Meaning           | Behavior                   |
|----------------|-------------------|----------------------------|
| `ExitContinue` | Normal execution  | Continue running tests     |
| `ExitOK`       | Command completed | Help or version printed    |
| `ExitUsage`    | CLI usage error   | Invalid flags or arguments |

---

## Flag Behavior

| Flag | Long Form         | Description                                    |
|------|-------------------|------------------------------------------------|
| `-c` | `--with-coverage` | Enables coverage profile generation            |
| `-o` | `--open-coverage` | Opens the coverage report after tests complete |
| `-q` | `--quiet`         | Suppresses verbose test output                 |
| `-V` | `--clean-view`    | Hides `[no test files]` lines from output      |
| `-I` | `--ignore-pattern`| Excludes packages matching the given pattern   |
| `-h` | `--help`          | Prints CLI usage information                   |
| `-v` | `--version`       | Prints the GoTestX version                     |

---

## Flag Interaction

| Flag              | Implicit Behavior                       |
|-------------------|-----------------------------------------|
| `--open-coverage` | Automatically enables `--with-coverage` |

Example:

```bash
gotestx -o
```

Equivalent to:

```bash
gotestx --with-coverage --open-coverage
```

---

## Default Package Resolution

If no packages are provided, GoTestX defaults to:

    ./...

Example:

```bash
gotestx
```

This runs tests across the entire module.

You may also specify packages manually:

```bash
gotestx ./internal ./cmd
```

---

## Combined Short Flags

Short flags may be combined in a single argument.

Example:

```bash
gotestx -cqV
```

Equivalent to:

```bash
gotestx -c -q -V
```

---

## Package Exclusion

GoTestX allows excluding packages from test execution using:

- CLI flag: `-I <pattern>`
- Ignore file: `.gotestxignore`

### Examples

Exclude all packages containing a `mock` segment:

```bash
gotestx -I mock
```

Exclude a specific path:

```bash
gotestx -I outport/testkit
```

Multiple exclusions:

```bash
gotestx -I mock -I testkit
```

### Ignore File

You can define persistent exclusions in a `.gotestxignore` file:

```
mock
outport/testkit
```

### Pattern Behavior

Patterns use a **tree-based syntax**:

| Pattern            | Matches                                  |
|--------------------|------------------------------------------|
| `mock`             | any path segment named `mock`             |
| `outport/testkit`  | exact subpath match                      |

Matching is segment-based (not substring-based).

### Execution Behavior

- Filtering occurs **after package discovery**
- If all packages are excluded, execution exits successfully:

```
No packages to test after applying ignore rules.
```

---

## Output Modes

### Quiet Mode

Quiet mode suppresses verbose output and prints a concise summary.

Example:

```bash
gotestx -q
```

Possible output:

    ✅ Tests finished successfully

With coverage enabled:

```bash
gotestx -cq
```

Example output:

    coverage: 92.3% of statements

### Clean View

Clean view removes noisy Go test output lines such as:

    ? pkg/name [no test files]

Example:

```bash
gotestx -V
```

---

## Coverage Workflow

```bash
gotestx -c
```

Generates:

    coverage.out

View:

```bash
go tool cover -html=coverage.out
```

Open automatically:

```bash
gotestx -co
```

Note: only supported on **macOS**.

---

## Internal Architecture

| File            | Responsibility                                  |
|-----------------|-------------------------------------------------|
| `options.go`    | CLI argument parsing                            |
| `run.go`        | Test execution orchestration                    |
| `coverage.go`   | Coverage helpers                               |
| `quiet.go`      | Quiet-mode output processing                    |
| `clean_view.go` | Output filtering                               |
| `ignore.go`     | Package exclusion logic                         |
| `ignorefile.go` | `.gotestxignore` loading                        |
| `runtime.go`    | Runtime utilities                              |
| `command.go`    | Command abstraction                            |
| `usage.go`      | CLI help                                       |
| `version.go`    | Version info                                   |
| `exitcodes.go`  | Exit semantics                                 |

---

## Design Goals

- keep CLI simple and predictable
- minimal dependencies
- deterministic testing
- isolated execution orchestration
- flexible package exclusion
- extensible architecture

---

## Package Visibility

The package is under `internal/` to prevent external usage.
