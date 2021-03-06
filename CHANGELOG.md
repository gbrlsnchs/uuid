# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [0.6.0] - 2018-10-17
### Added
- UUID builders that don't return an `error`, but panics.
- `sql.Scanner` and `driver.Value` implementations.
- Test cases for `UUID.GUID` method.
- Test cases for marshaling methods.
- Test cases for unmarshaling methods.
- Test cases for `UUID.Scan`.

### Changed
- Accept an alternative `io.Reader` when building a version 4 UUID.

### Fixed
- Fix `UUID.UnmarshalBinary` not properly setting the unmarshaled data to the receiver.

### Removed
- `ParseBytes` function (it is now simply `Parse`).

## [0.5.0] - 2018-09-03
### Added
- Sourcegraph badge.

### Changed
- Rename `ErrInvalidUUID` to `ErrInvalid`.

## [0.4.0] - 2018-09-03
### Changed
- Export `ParseBytes`.

## [0.3.0] - 2018-09-03
### Added
- AppVeyor config and badge.

### Changed
- Change UUID builders' name prefixes from `Create` to `Generate`.

## [0.2.0] - 2018-08-31
### Changed
- Allow random bits to be used instead of a MAC address for V1 and V2.
- Use different regexps for each version in tests.

## 0.1.0 - 2018-08-31
### Added
- This changelog file.
- README file.
- MIT License.
- Travis CI configuration file.
- EditorConfig file.
- `go.mod` file.
- Support for all UUID versions.
- Tests.

[0.6.0]: https://github.com/gbrlsnchs/uuid/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/gbrlsnchs/uuid/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/gbrlsnchs/uuid/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/gbrlsnchs/uuid/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/gbrlsnchs/uuid/compare/v0.1.0...v0.2.0
