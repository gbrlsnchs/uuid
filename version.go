package uuid

// Version is a modifying byte.
type Version byte

const (
	// VersionNone is an invalid UUID version.
	VersionNone Version = iota
	// Version1 is the version 1 of the UUID implementation.
	Version1 Version = iota << 4
	// Version2 is the version 2 of the UUID implementation.
	Version2
	// Version3 is the version 3 of the UUID implementation.
	Version3
	// Version4 is the version 4 of the UUID implementation.
	Version4
	// Version5 is the version 5 of the UUID implementation.
	Version5
	// versionByte is the index of the version byte.
	versionByte byte = 6
)

func (v Version) String() string {
	switch v {
	case Version1:
		return "Version 1"
	case Version2:
		return "Version 2"
	case Version3:
		return "Version 3"
	case Version4:
		return "Version 4"
	case Version5:
		return "Version 5"
	default:
		return "No version"
	}
}
