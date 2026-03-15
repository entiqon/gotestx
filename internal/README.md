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
| `-h` | `--help`          | Prints CLI usage information                   |
| `-v` | `--version`       | Prints the GoTestX version                     |

---

## Flag Interaction

| Flag              | Implicit Behavior                       |
|-------------------|-----------------------------------------|
| `--open-coverage` | Automatically enables `--with-coverage` |

Example:

``` bash
gotestx -o
```

Equivalent to:

``` bash
gotestx --with-coverage --open-coverage
```

---

## Default Package Resolution

If no packages are provided, GoTestX defaults to:

    ./...

Example:

``` bash
gotestx
```

This runs tests across the entire module.

You may also specify packages manually:

``` bash
gotestx ./internal ./cmd
```

---

## Combined Short Flags

Short flags may be combined in a single argument.

Example:

``` bash
gotestx -cqV
```

Equivalent to:

``` bash
gotestx -c -q -V
```

---

## Output Modes

### Quiet Mode

Quiet mode suppresses verbose output and prints a concise summary.

Example:

``` bash
gotestx -q
```

Possible output:

    ✅ Tests finished successfully

With coverage enabled:

``` bash
gotestx -cq
```

Example output:

    coverage: 92.3% of statements

### Clean View

Clean view removes noisy Go test output lines such as:

    ? pkg/name [no test files]

Example:

``` bash
gotestx -V
```

This makes test output easier to read when many packages contain no test
files.

---

## Coverage Workflow

Running GoTestX with coverage:

``` bash
gotestx -c
```

Generates a coverage profile:

    coverage.out

You can view the coverage report using:

``` bash
go tool cover -html=coverage.out
```

To automatically open the report after running tests:

``` bash
gotestx -co
```

Note: opening the coverage report is currently supported on **macOS**.

---

## Internal Architecture

| File            | Responsibility                                  |
|-----------------|-------------------------------------------------|
| `options.go`    | CLI argument parsing                            |
| `run.go`        | Test execution orchestration                    |
| `coverage.go`   | Coverage argument generation and report opening |
| `quiet.go`      | Quiet-mode output processing                    |
| `clean_view.go` | Filtering noisy test output                     |
| `runtime.go`    | Runtime utilities (OS helpers)                  |
| `command.go`    | Command execution abstraction                   |
| `usage.go`      | CLI help output                                 |
| `version.go`    | Version information                             |
| `exclusions.go` | Output filtering rules                          |
| `exitcodes.go`  | CLI exit code definitions                       |

---

## Execution Flow

The typical execution flow of GoTestX is:

    main.go
       ↓
    ResolveOptions
       ↓
    Run
       ↓
    command execution
       ↓
    test output processing
       ↓
    coverage handling

The `run.go` file acts as the **execution orchestrator**, delegating
specific responsibilities to smaller components such as:

- coverage helpers
- quiet output processing
- clean view filtering
- runtime utilities

This separation keeps the runtime easier to test and maintain.

---

## Design Goals

The internal runtime is designed to:

- keep the CLI simple and predictable
- maintain minimal dependencies
- allow deterministic testing
- keep test orchestration isolated from the CLI entrypoint
- support future extensibility of test execution workflows
- isolate CLI behavior from Go runtime implementation details

---

## Package Visibility

The package is placed under `internal/` intentionally to enforce Go
module visibility rules and prevent external consumers from depending on
unstable implementation details.

Only the CLI entrypoint (`main.go`) should interact with this package.
