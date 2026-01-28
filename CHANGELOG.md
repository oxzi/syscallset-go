<!--
SPDX-FileCopyrightText: Alvar Penning

SPDX-License-Identifier: BSD-3-Clause
-->

# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog][keep-a-changelog], and this project adheres to [Semantic Versioning][semantic-versioning].

## [Unreleased]
## [0.1.7] - 2026-01-28
### Changed
- LimitTo: Ignore unknown system calls as they might be relevant on other architectures.
- syscalls: bump systemd 256 -> 259
- Bump Go and CI dependencies

## [0.1.6] - 2024-11-29
### Changed
- syscalls: bump systemd 254 -> 256
- Bump Go and CI dependencies

## [0.1.5] - 2023-09-15
### Added
- generator/gen.go: syscall group description with auto-generated package documentation
- Write CHANGELOG for all prior changes.

### Changed
- syscalls: bump systemd 252 -> 254
- Bump Go and CI dependencies
- Remove year in copyright headers

### Fixed
- generator/gen.go: fix for systemd 254's new format

## [0.1.4] - 2022-12-04
### Changed
- update module github.com/elastic/go-seccomp-bpf to v1.3.0

## [0.1.3] - 2022-11-19
### Added
- Configure Renovate

### Changed
- syscalls: bump systemd 251 -> 252
- Bump Go and CI dependencies

## [0.1.2] - 2022-11-02
### Fixed
- generator: prevent duplicates and sort syscalls

## [0.1.1] - 2022-10-08
### Changed
- syscalls: bump systemd 250 -> 251

## [0.1.0] - 2022-03-10
### Added
- First release with working code.

[keep-a-changelog]: https://keepachangelog.com/en/1.1.0/
[semantic-versioning]: https://semver.org/spec/v2.0.0.html

[unreleased]: https://github.com/oxzi/syscallset-go/compare/v0.1.7..HEAD
[0.1.7]: https://github.com/oxzi/syscallset-go/compare/v0.1.6..v0.1.7
[0.1.6]: https://github.com/oxzi/syscallset-go/compare/v0.1.5..v0.1.6
[0.1.5]: https://github.com/oxzi/syscallset-go/compare/v0.1.4..v0.1.5
[0.1.4]: https://github.com/oxzi/syscallset-go/compare/v0.1.3..v0.1.4
[0.1.3]: https://github.com/oxzi/syscallset-go/compare/v0.1.2..v0.1.3
[0.1.2]: https://github.com/oxzi/syscallset-go/compare/v0.1.1..v0.1.2
[0.1.1]: https://github.com/oxzi/syscallset-go/compare/v0.1.0..v0.1.1
[0.1.0]: https://github.com/oxzi/syscallset-go/releases/tag/v0.1.0
