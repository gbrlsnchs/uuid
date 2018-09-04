# uuid (Universally Unique IDentifier generator for Go)
[![Build status](https://travis-ci.org/gbrlsnchs/uuid.svg?branch=master)](https://travis-ci.org/gbrlsnchs/uuid)
[![Build status](https://ci.appveyor.com/api/projects/status/ofys86q2b22b4rlk/branch/master?svg=true)](https://ci.appveyor.com/project/gbrlsnchs/uuid/branch/master)
[![Sourcegraph](https://sourcegraph.com/github.com/gbrlsnchs/uuid/-/badge.svg)](https://sourcegraph.com/github.com/gbrlsnchs/uuid?badge)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/uuid?status.svg)](https://godoc.org/github.com/gbrlsnchs/uuid)
[![Minimal version](https://img.shields.io/badge/minimal%20version-go1.11%2B-5272b4.svg)](https://golang.org/doc/go1.11)

## About
This package is a UUID (or GUID) generator for [Go](https://golang.org). It's still under development.

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
guid, err := uuid.GenerateV4()
if err != nil {
	// handle error
}
log.Print(guid.String())           // prints a 36-byte hex-encoded UUID
log.Print(guid.Version().String()) // prints "Version 4"
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
log.Print(guid.Version().String()) // prints "Version 4"
log.Print(guid.Variant().String()) // prints "RFC 4122"
```

## Contributing
### How to help
- Pull Requests
- Issues
- Opinions
