# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [1.1.1] - 2026-03-15

### Added
- Governance and community documentation:
  - `CODEOWNERS`
  - `CONTRIBUTING.md`
  - `CONTRIBUTORS.md`
  - `SECURITY.md`
  - `SUPPORT.md`
  - `CODE_OF_CONDUCT.md`
  - `NOTICE`
  - `ROADMAP.md`
  - `RELEASE.md`

### Changed
- Established **canonical project attribution** and official repository under **Entiqon Labs**.
- Added standard copyright headers to core source files.

### Security
- Introduced repository protection rules:
  - required pull requests
  - required approvals
  - required signed commits
  - linear history enforcement
  - force-push prevention

---

## [1.1.0] - 2025-09-21

### Added
- **Quiet mode summary output**: `-q` now always prints a concise result:
  - ✅ success if all tests passed
  - coverage % if `-c` is enabled
  - ❌ failure with hint to rerun without `-q`
- **Sample Output section** in README for clearer documentation.

### Changed
- **`-C` flag renamed to `-V` (clean-view)** to avoid confusion with `-c` (coverage).
  - `-V` removes `[no test files]` lines from output.
  - Works in combination with `-q` and `-c`.
- Quiet mode errors (invalid path, no Go files, test failures) are now unified:
  ```
  ❌ Tests failed (use without -q to see details)
  ```

### Fixed
- Improved test coverage:
  - Covered quiet mode failure branch.
  - Covered default `./...` package expansion.
  - Covered ellipsis (`./...`) handling in arguments.

## [1.0.0] - 2025-09-18

### Added
- Initial release of **GoTestX**.
- Extended `go test` with:
  - Coverage reporting (`-c`).
  - Open coverage in browser on macOS (`-o`).
  - Quiet mode (`-q`).
  - Clean output (`-C`).
- Support for combined flags (e.g. `-cq`, `-coq`, `-cC`).
- Smart package detection and validation.
