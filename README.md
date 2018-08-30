# uuid (Universally Unique IDentifier generator for Go)
[![Build Status](https://travis-ci.org/gbrlsnchs/uuid.svg?branch=master)](https://travis-ci.org/gbrlsnchs/uuid)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/uuid?status.svg)](https://godoc.org/github.com/gbrlsnchs/uuid)

## About
This package is a UUID (or GUID) generator for [Go]. It's still under development.

### Supported versions:
| Version | Supported          |
|:-------:|:------------------:|
| 1       | :heavy_check_mark: |
| 2       | :heavy_check_mark: |
| 3       | :heavy_check_mark: |
| 4       | :heavy_check_mark: |
| 5       | :heavy_check_mark: |

## Usage
Full documentation [here] (work in progress).

## Example
### Generating UUIDs
```go
guid, err := uuid.V4()
if err != nil {
	// handle error
}
log.Print(guid.String())           // prints a 36-byte hex-encoded UUID
log.Print(guid.Version().String()) // prints "V4"
log.Print(guid.Variant().String()) // prints "RFC 4122"
```

### Building UUIDs from strings
```go
guid, err := uuid.Parse("d9ab3f01-482f-425d-8a10-a24b0abfe661")
if err != nil {
	// handle error
}
log.Print(guid.String())           // prints "d9ab3f01-482f-425d-8a10-a24b0abfe661"
log.Print(guid.GUID())             // prints "{d9ab3f01-482f-425d-8a10-a24b0abfe661}"
log.Print(guid.Version().String()) // prints "V4"
log.Print(guid.Variant().String()) // prints "RFC 4122"
```

## Contribution
### How to help:
- Pull Requests
- Issues
- Opinions

[Go]: https://golang.org
[here]: https://godoc.org/github.com/gbrlsnchs/uuid
