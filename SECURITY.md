# Security Policy

## Supported Versions

The GoTestX project maintains security fixes for the most recent releases.

| Version                | Supported       |
|------------------------|-----------------|
| Latest release         | ✅               |
| Previous minor release | ⚠️ Best effort  |
| Older versions         | ❌ Not supported |

Users are strongly encouraged to upgrade to the latest release to receive security fixes.

---

# Reporting a Vulnerability

If you discover a security vulnerability in GoTestX, please report it responsibly.

Do **not** open a public GitHub issue for security vulnerabilities.

Instead, report the issue privately to the project maintainer.

**Contact**

Author / Maintainer:
Isidro A. López G.

Project:
https://github.com/entiqon/gotestx

Please include the following information when reporting a vulnerability:

* A description of the issue
* Steps to reproduce the vulnerability
* Potential impact
* Suggested mitigation if known

We will acknowledge receipt of your report as soon as possible and work to address the issue promptly.

---

# Security Considerations

GoTestX is a CLI wrapper around the Go toolchain and executes the `go test` command internally.

The project takes the following precautions:

* Arguments passed to the Go toolchain are not executed through a shell.
* Command execution is handled using Go's `exec.Command`.
* No external dependencies are required by the core runtime.
* The tool does not execute arbitrary scripts or external binaries beyond the Go toolchain.

However, users should
