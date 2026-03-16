# Changelog
All notable changes to this project will be documented in this file.

### Changed

- **core**: Extract coverage, quiet, and clean view helpers and refine CLI runtime infrastructure


### build

- **release**: Enforce releases from main branch


### security

- Update README with official release verification notice


### Documentation

- **release**: Add releases section and git-cliff configuration


### Fixed

- **release**: Remove unsupported --template flag from git-cliff invocation

- **release**: Stop tracking generated release notes


### build

- **release**: Move release pipeline to Makefile and remove CI workflow

- **release**: Add full release target and ignore dist artifacts


### chore

- **build**: Introduce Makefile for build, test, install, and release tasks


### ci

- **release**: Introduce automated GitHub release workflow with git-cliff and binary builds

- **release**: Automate changelog generation and GitHub releases using git-cliff

- **changelog**: Format commit scopes as bold package identifiers

- **release**: Add git-cliff release template


### Documentation

- Establish canonical project attribution

- Establish canonical project attribution


### ci

- **gotestx**: Typo when download gotestx tool


### Documentation

- **readme**: Update for standalone GoTestX repository


### Fixed

- Correct gotestx tool url


### chore

- **release**: Bump GoTestX to v1.1.0 with quiet mode updates


### Added

- Initial import of GoTestX v1.0.0

