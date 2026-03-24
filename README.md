# GoTestX

> A modern, developer-friendly test runner for Go with coverage, filtering, and clean output.

[![Official Repository](https://img.shields.io/badge/Repository-Entiqon%20Labs-blue?logo=github)](https://github.com/entiqon/gotestx)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/entiqon/gotestx)](https://github.com/entiqon/gotestx/releases)
[![Build Status](https://github.com/entiqon/gotestx/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/gotestx/actions)
[![Codecov](https://codecov.io/gh/entiqon/gotestx/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/gotestx)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/gotestx)](https://goreportcard.com/report/github.com/entiqon/gotestx)
[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/gotestx.svg)](https://pkg.go.dev/github.com/entiqon/gotestx)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Security Verified](https://img.shields.io/badge/Security-Verified-success?logoColor=white)](https://github.com/Entiqon/gotestx/security/advisories)
[![GPG Signed](https://img.shields.io/badge/GPG-Signed-blue?logo=gnupg&logoColor=white)](https://github.com/Entiqon/gotestx/commits/main)

**Official repository (canonical source):** https://github.com/entiqon/gotestx

> [!CAUTION]
> **OFFICIAL INSTALLATION ONLY**
> Please ensure you are using the official **Entiqon/GoTestX** repository.
> Never download binaries from untrusted forks. Use only the official releases.

---

## Get started

GoTestX is a developer-focused test runner for Go that simplifies and enhances
the standard `go test` workflow.

It provides a consistent CLI with built-in support for coverage, package exclusion,
quiet execution, and clean output — without requiring additional scripting.

---

## ✨ Features

* **Coverage mode** (`-c`): generates `coverage.out` using `-covermode=atomic`.
* **Open coverage** (`-o`): opens the HTML coverage report (macOS only).
* **Package exclusion** (`-I`): exclude packages using patterns or `.gotestxignore`.
* **Quiet mode** (`-q`): minimal output with clear success/failure signals.
* **Clean view** (`-V`): removes noisy `[no test files]` lines.
* **Composable flags**: combine short flags (e.g. `-cq`, `-coq`, `-cVq`).
* **Smart package resolution**:
  * Expands `./pkg` → `./pkg/...` when appropriate.
  * Validates paths and reports clear errors.

---

## 🚀 Installation

```bash
go install github.com/entiqon/gotestx@latest
```

Verify:

```bash
gotestx -v
```

---

## 📦 Usage

```bash
gotestx [flags] [packages]
```

### Flags

```
-c, --with-coverage     Enable coverage profile generation (coverage.out)
-o, --open-coverage     Open coverage report (macOS only, implies -c)
-q, --quiet             Suppress verbose output (summary only)
-V, --clean-view        Hide '[no test files]' lines
-I, --ignore-pattern    Exclude packages matching the given pattern
-h, --help              Show help
-v, --version           Show version info
```

---

## 🚫 Package Exclusion

Exclude packages using CLI flags or `.gotestxignore`.

### Examples

```bash
gotestx -I mock ./...
gotestx -I mock -I testkit ./...
```

### Ignore file

```
mock
outport/testkit
```

### Pattern behavior

| Pattern           | Matches                      |
|-------------------|-----------------------------|
| mock              | segment named "mock"         |
| outport/testkit   | exact subpath               |

### Behavior

- Applied after package discovery
- If all packages are excluded:

```
No packages to test after applying ignore rules.
```

---

## 🧪 Examples

```bash
gotestx
gotestx -c ./...
gotestx -cq ./...
gotestx -co ./...
gotestx -V ./...
gotestx -cVq ./...
```

---

## 🖥 Sample Output

```
Running tests normally across: ./internal/...
ok  	github.com/entiqon/gotestx/internal	0.359s
```

Quiet:

```
✅ Tests finished successfully
```

Failure:

```
❌ Tests failed (use without -q to see details)
```

---

## 🛠 Development

```bash
git clone https://github.com/entiqon/gotestx.git
cd gotestx
go build -o gotestx .
go test ./internal/... -v
```

---

## 📄 License

MIT License — Entiqon Labs
