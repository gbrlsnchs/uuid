# uuid (Universally Unique IDentifier generator for Go)
[![Build Status](https://travis-ci.org/gbrlsnchs/uuid.svg?branch=master)](https://travis-ci.org/gbrlsnchs/uuid)
[![Sourcegraph](https://sourcegraph.com/github.com/gbrlsnchs/uuid/-/badge.svg)](https://sourcegraph.com/github.com/gbrlsnchs/uuid?badge)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/uuid?status.svg)](https://godoc.org/github.com/gbrlsnchs/uuid)
[![Minimal Version](https://img.shields.io/badge/minimal%20version-go1.10%2B-5272b4.svg)](https://golang.org/doc/go1.10)

## About
This package is a UUID (or GUID) generator for [Go](https://golang.org).

### Supported versions:
| Version | Supported          |
|:-------:|:------------------:|
| 1       | :heavy_check_mark: |
| 2       | :heavy_check_mark: |
| 3       | :heavy_check_mark: |
| 4       | :heavy_check_mark: |
| 5       | :heavy_check_mark: |

## Usage
Full documentation [here](https://godoc.org/github.com/gbrlsnchs/uuid).

### Installing
#### Go 1.10
`vgo get -u github.com/gbrlsnchs/uuid`
#### Go 1.11 or after
`go get -u github.com/gbrlsnchs/uuid`

### Importing
```go
import (
	// ...

	"github.com/gbrlsnchs/uuid"
)
```

## Example
### Generating UUIDs
```go
guid := uuid.V4(nil)                            // panics if there's an error
log.Printf("guid = %v", guid)                   // prints a 36-byte hex-encoded UUID
log.Printf("guid version = %v", guid.Version()) // prints "Version 4"
log.Printf("guid variant = %v", guid.Variant()) // prints "RFC 4122"
```

### Building UUIDs from strings
```go
guid, err := uuid.Parse("d9ab3f01-482f-425d-8a10-a24b0abfe661")
if err != nil {
	// handle error
}
log.Print(guid.String())           // prints "d9ab3f01-482f-425d-8a10-a24b0abfe661"
log.Print(guid.GUID())             // prints "{d9ab3f01-482f-425d-8a10-a24b0abfe661}"
log.Print(guid.Version().String()) // prints "Version 4"
log.Print(guid.Variant().String()) // prints "RFC 4122"
```

## Contributing
### How to help
- For bugs and opinions, please [open an issue](https://github.com/gbrlsnchs/uuid/issues/new)
- For pushing changes, please [open a pull request](https://github.com/gbrlsnchs/uuid/compare)
