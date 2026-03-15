# Release Process

This document describes the release process for **GoTestX**.

GoTestX follows a lightweight release workflow designed for CLI tools.

---

# Versioning

GoTestX follows **Semantic Versioning (SemVer)**.

Format:

```
MAJOR.MINOR.PATCH
```

Example:

```
1.1.0
```

Meaning:

| Type  | When to use                      |
| ----- | -------------------------------- |
| MAJOR | Breaking CLI or API changes      |
| MINOR | New features                     |
| PATCH | Bug fixes and small improvements |

---

# Preparing a Release

Before creating a new release, ensure the following:

1. All tests pass

```
go test ./...
```

2. Static analysis passes

```
go vet ./...
```

3. Code is formatted

```
go fmt ./...
```

4. Documentation is updated if necessary

* README.md
* ROADMAP.md
* CLI help text

---

# Updating the Version

Update the version constant inside the source code.

Example:

```go
const Version = "1.2.0"
```

---

# Creating a Release

Create a Git tag for the new version.

Example:

```
git tag v1.2.0
git push origin v1.2.0
```

GitHub Actions will automatically run the CI pipeline.

---

# Publishing a GitHub Release

After pushing the tag:

1. Go to the GitHub repository
2. Navigate to **Releases**
3. Click **Draft a new release**
4. Select the tag

Add release notes describing:

* New features
* Improvements
* Bug fixes

Example:

```
## v1.2.0

Features
- Added coverage summary flag (-s)

Improvements
- Improved quiet mode output

Fixes
- Corrected clean-view filtering behavior
```

---

# Installation

Users install GoTestX via:

```
go install github.com/entiqon/gotestx@latest
```

Or install a specific version:

```
go install github.com/entiqon/gotestx@v1.2.0
```

---

# Verifying the Release

After installation:

```
gotestx -v
```

Expected output example:

```
1.2.0

GoTestX
Go Test eXtended tool with coverage support
Author: Entiqon Project Team
Version: 1.2.0
```

---

# Release Guidelines

Maintain the following principles:

* Releases should be small and incremental
* Avoid breaking CLI behavior
* Document all visible changes
* Ensure tests pass before tagging

---

# Maintainer

* Original Author: `Isidro A. López G.`
* Lead Architect:
  * `Isidro A. López G.`
  * `Gailisis Dawsons`
* Organization: `Entiqon Labs`
* Official repository: `https://github.com/entiqon/gotestx`
