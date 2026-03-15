# GoTestX

> Go Test eXtended tool with coverage support

[![Official Repository](https://img.shields.io/badge/Repository-Entiqon%20Labs-blue?logo=github)](https://github.com/entiqon/gotestx)
[![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![Release](https://img.shields.io/github/v/release/entiqon/gotestx)](https://github.com/entiqon/gotestx/releases)
[![Build Status](https://github.com/entiqon/gotestx/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/gotestx/actions)
[![Codecov](https://codecov.io/gh/entiqon/gotestx/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/gotestx)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/gotestx)](https://goreportcard.com/report/github.com/entiqon/gotestx)
[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/gotestx.svg)](https://pkg.go.dev/github.com/entiqon/gotestx)
[![License](https://img.shields.io/github/license/entiqon/gotestx)](LICENSE)

**Official repository (canonical source):** https://github.com/entiqon/gotestx

## Get started

GoTestX extends the standard [`go test`](https://pkg.go.dev/cmd/go#hdr-Test_packages) command with a simpler, more
versatile interface.  
It adds optional coverage reporting, quiet mode, and clean output filtering — while remaining fully compatible with
`go test`.

---

## ✨ Features

* **Coverage mode** (`-c`): generates `coverage.out` with `-covermode=atomic`.
* **Open coverage** (`-o`): opens the HTML coverage report in a browser (macOS only).
* **Quiet mode** (`-q`): suppresses verbose chatter, but always reports:
    * ✅ success if all tests passed
    * coverage % if `-c` is enabled
    * ❌ failure (with hint to rerun without `-q`)
* **Clean view** (`-V`): removes `? … [no test files]` lines for cleaner output.
* **Flag combinations**: short flags can be combined (e.g. `-cq`, `-coq`, `-cVq`).
* **Smart package detection**:
    * Expands `./pkg` → `./pkg/...` if root has no Go files but subpackages do.
    * Reports errors if a path doesn’t exist or has no Go files.

---

## 🚀 Installation

Recommended:

``` bash
go install github.com/entiqon/gotestx@latest
```

Check installation:

``` bash
gotestx -v
```

---

## 📦 Releases

Prebuilt binaries are available for users who prefer not to build the
tool locally.

Latest version:

https://github.com/entiqon/gotestx/releases

Each release includes binaries for:

- Linux (amd64)
- macOS Intel (amd64)
- macOS Apple Silicon (arm64)
- Windows (amd64)

Example installation from binary:

``` bash
chmod +x gotestx
sudo mv gotestx /usr/local/bin/
```

However, installing via Go remains the recommended approach.

------------------------------------------------------------------------

## 📦 Usage

``` bash
gotestx [options] [packages]
```

Options:

```
  -c, --with-coverage   Run tests with coverage report generation (coverage.out)
  -o, --open-coverage   Open coverage report in browser (macOS only, implies -c)
  -q, --quiet           Suppress verbose output, only show summary/coverage/fail
  -V, --clean-view      Suppress 'no test files' lines for cleaner output
  -h, --help            Show this help
  -v, --version         Show version info
```

---

## 🧪 Examples

Run tests for all packages:

```bash
gotestx
```

Run tests with coverage:

```bash
gotestx -c ./...
```

Run quietly with coverage (only one summary line):

```bash
gotestx -cq ./...
```

Run with coverage and open report in browser (macOS):

```bash
gotestx -o ./...
```

Run with clean output (no `[no test files]` lines):

```bash
gotestx -V ./...
```

Combine flags:

```bash
gotestx -cVq ./...
```

---

## 🖥 Sample Output

Normal run:

```
Running tests normally across: ./internal/...
ok  	github.com/entiqon/gotestx/internal	0.359s
```

Quiet run:

```
✅ Tests finished successfully
```

Quiet + coverage:

```
ok  	github.com/entiqon/gotestx/internal	0.359s	coverage: 100.0% of statements
```

Quiet with failure:

```
❌ Tests failed (use without -q to see details)
```

Clean view:

```
ok  	github.com/entiqon/gotestx/internal/join	0.01s
```

(no `[no test files]` lines)

---

## 🛠 Development

Clone the repository:

```bash
git clone https://github.com/entiqon/gotestx.git
cd gotestx
```

Build:

```bash
go build -o gotestx .
```

Test:

```bash
go test ./internal/... -v
```

---

## 🏛️ Ownership & Attribution

**GoTestX** is an original project developed and maintained by **Entiqon Labs**.

* **Lead Architect:** [Isidro A. López G.](https://github.com/ialopezg)
* **Official Repository:** [github.com/entiqon/gotestx](https://github.com/entiqon/gotestx)
* **Initial Release:** September 2025

---

## 📄 License

Part of the [Entiqon Project](https://github.com/entiqon).  
Licensed under the **MIT License**. As per the license terms, the original copyright notice and this permission notice
must be included in all copies or substantial portions of the software.
