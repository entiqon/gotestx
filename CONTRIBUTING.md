# Contributing to GoTestX

Thank you for your interest in contributing to **GoTestX**.

This project is part of the **Entiqon Labs developer tooling ecosystem** and is maintained by the **Entiqon Labs Team**.

* **Lead Architects:** Entiqon Labs Team
* **Organization:** Entiqon Labs
* **Official Repository:** https://github.com/entiqon/gotestx

---

# Attribution & Intellectual Ownership

GoTestX was originally designed and implemented by **Isidro A. López G.** under the **Entiqon Project**.

All contributions must respect the following:

* The original copyright notice **must not be removed**
* The original author attribution **must remain intact**
* Derivative work must reference the **original repository**

The MIT License permits reuse, modification, and redistribution, but **proper attribution is mandatory**.

If you reuse or fork this project, you must include:

```
Original project: GoTestX
Author: Isidro A. López G.
Repository: https://github.com/entiqon/gotestx
```

Failure to preserve attribution violates the license terms.

---

# Code of Contribution

Contributors should aim to:

* Keep GoTestX **minimal**
* Avoid unnecessary abstractions
* Maintain compatibility with `go test`
* Preserve the **testable CLI architecture**

Every change should improve developer productivity without turning GoTestX into a heavy framework.

---

# Development Setup

Clone the repository:

```
git clone https://github.com/entiqon/gotestx.git
cd gotestx
```

Build the binary:

```
go build -o gotestx .
```

Run tests:

```
go test ./internal/... -v
```

---

# Project Structure

```
gotestx/
    main.go
    internal/
        run.go
        flags.go
        coverage.go
```

Core logic lives inside `internal`.

The CLI entry point is intentionally minimal.

---

# Coding Guidelines

Please follow standard Go conventions.

### Formatting

```
go fmt ./...
```

### Static analysis

```
go vet ./...
```

### Tests

All new features must include tests when possible.

The project uses a **testable CLI pattern**:

```
Run(args, stdout, stderr)
```

This allows CLI behavior to be tested without spawning subprocesses.

---

# Commit Guidelines

Use conventional commit style:

```
feat: add coverage summary support
fix: correct quiet mode output handling
docs: update README examples
test: add CLI flag parsing tests
```

---

# Pull Request Process

1. Fork the repository
2. Create a feature branch
3. Implement changes
4. Add tests if applicable
5. Submit a Pull Request

PRs should include:

* Clear description
* Motivation for the change
* Example usage if applicable

---

# What Not to Contribute

Please avoid submitting PRs that introduce:

* unnecessary dependencies
* heavy frameworks
* features unrelated to Go testing
* breaking CLI behavior

GoTestX intentionally remains **small and focused**.

---

# License

By contributing to this project you agree that your contributions will be licensed under the **MIT License**.

See the `LICENSE` file for details.

---

# Maintainers

Lead Architects: **Entiqon Labs Team**

Organization: **Entiqon Labs**
