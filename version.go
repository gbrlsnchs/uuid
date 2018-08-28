package uuid

// Version is a modifying byte.
type Version byte

const (
	// Version1 is the version 1 of the UUID implementation.
	Version1 Version = (iota + 1) << 4
	// Version2 is the version 2 of the UUID implementation.
	Version2
	// Version3 is the version 3 of the UUID implementation.
	Version3
	// Version4 is the version 4 of the UUID implementation.
	Version4
	// Version5 is the version 5 of the UUID implementation.
	Version5
	// versionByte is the index of the version byte.
	versionByte = 6
)

func (v Version) String() string {
	switch v {
	case Version1:
		return "V1"
	case Version2:
		return "V2"
	case Version3:
		return "V3"
	case Version4:
		return "V4"
	default:
		return "V5"
	}
}
