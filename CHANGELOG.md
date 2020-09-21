# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Added sidecar container to setup vault's kubernetes auth backend.

## [2.0.0] - 2020-08-11

### Changed

- Updated Kubernetes dependencies to v1.18.5.
- Updated helmclient to 2.0.0.

## [0.3.0] 2020-05-19

### Changed

- Set node port in vault helm chart.
- Push Helm Chart to Appr on Quay.
- Upgrade helmclient to use Helm 3.

## [0.2.0] 2020-03-26

### Changed

- Switched from dep to go modules.
- Switched to architect-orb.

## [0.1.0] 2020-03-19

### Added

- First release.

[Unreleased]: https://github.com/giantswarm/e2esetup/compare/v2.0.0...HEAD
[2.0.0]: https://github.com/giantswarm/e2esetup/compare/v0.3.0...v2.0.0
[0.3.0]: https://github.com/giantswarm/e2esetup/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/e2esetup/compare/v0.1.0...v0.2.0

[0.1.0]: https://github.com/giantswarm/e2esetup/releases/tag/v0.1.0
