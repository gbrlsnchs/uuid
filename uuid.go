package uuid

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"unsafe"
)

const (
	byteSize  = 16
	hexSize   = byteSize*2 + 4 // double the bytes + 4 bytes of dashes
	urnPrefix = "urn:uuid:"
	urnSize   = hexSize + len(urnPrefix)
	urnOffset = urnSize - hexSize
	guidSize  = hexSize + 2
)

var (
	// Null is a null UUID that translates to a 36-byte string with only zeroes.
	Null = UUID{}
	// ErrInvalidUUID is the error for invalid UUID formats.
	ErrInvalidUUID = errors.New("uuid: invalid uuid")
	// ErrUnsupportedVersion is the error for non-existent or not implemented UUID versions.
	ErrUnsupportedVersion = errors.New("uuid: unsupported version")
)

// UUID is a 16-byte array Universally Unique IDentifier, as per the RFC-4122.
type UUID [byteSize]byte

// FromString parses a UUID 36-byte string encoded in hexadecimal and converts it to a 16-byte array.
func FromString(s string) (UUID, error) {
	return parseBytes([]byte(s))
}

// Generate generates a UUID (or GUID) according to the RFC-4122.
func Generate(v Version) (guid UUID, err error) {
	switch v {
	case Version4:
		guid, err = generateV4()
	default:
		return Null, ErrUnsupportedVersion
	}
	if err != nil {
		return Null, err
	}
	guid[versionByte] = guid[versionByte]&0x0F | byte(v)          // (byte & 00001111) | 01000000
	guid[variantByte] = guid[variantByte]&0x3F | byte(VariantRFC) // (byte & 00111111) | 10000000
	return guid, nil
}

func parseBytes(b []byte) (UUID, error) {
	if len(b) != hexSize {
		if len(b) != urnSize {
			return Null, ErrInvalidUUID
		}
		if !bytes.HasPrefix(b, []byte(urnPrefix)) {
			return Null, ErrInvalidUUID
		}
		b = b[urnOffset:]
	}
	var (
		guid UUID
		err  error
	)
	if _, err = hex.Decode(guid[:4], b[:8]); err != nil {
		return Null, err
	}
	if _, err = hex.Decode(guid[4:6], b[9:13]); err != nil {
		return Null, err
	}
	if _, err = hex.Decode(guid[6:8], b[14:18]); err != nil {
		return Null, err
	}
	if _, err = hex.Decode(guid[8:10], b[19:23]); err != nil {
		return Null, err
	}
	if _, err = hex.Decode(guid[10:], b[24:]); err != nil {
		return Null, err
	}
	return guid, nil
}

// Bytes returns the UUID as a 16-byte slice.
func (guid UUID) Bytes() []byte {
	return guid[:]
}

// GUID returns a 36-byte string with surrounding curly braces.
func (guid UUID) GUID() string {
	var microsoft [guidSize]byte
	last := len(microsoft) - 1
	microsoft[0] = '{'
	microsoft[last] = '}'
	guid.encode(microsoft[1:last])
	return guid.str(microsoft[:])
}

// IsNull returns whether the UUID is a null UUID.
func (guid UUID) IsNull() bool {
	return guid == Null
}

// MarshalBinary implements binary marshalling.
func (guid UUID) MarshalBinary() ([]byte, error) {
	return guid.Bytes(), nil // 16-byte array
}

// MarshalJSON implements JSON marshalling.
func (guid UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(guid.String()) // 36-byte hex-encoded string
}

// MarshalText implements text marshalling.
func (guid UUID) MarshalText() ([]byte, error) {
	var xid [hexSize]byte
	b := xid[:]
	guid.encode(b)
	return b, nil // 36-byte hex-encoded slice
}

// String converts the 16-byte UUID to a 36-byte string encoded in hexadecimal.
func (guid UUID) String() string {
	b, _ := guid.MarshalText()
	return guid.str(b)
}

// UnmarshalBinary implements binary unmarshalling.
func (guid *UUID) UnmarshalBinary(b []byte) error {
	if len(b) != byteSize {
		return ErrInvalidUUID
	}
	copy(guid.Bytes(), b)
	return nil
}

// UnmarshalJSON implements JSON unmarshalling.
func (guid *UUID) UnmarshalJSON(b []byte) error {
	return guid.UnmarshalText(b)
}

// UnmarshalText implements text unmarshalling.
func (guid *UUID) UnmarshalText(b []byte) error {
	u, err := parseBytes(b)
	if err != nil {
		return err
	}
	*guid = u
	return nil
}

// URN returns the UUID as string conformed to the RFC-2141.
func (guid UUID) URN() string {
	var urn [urnSize]byte
	b := urn[:]
	copy(b, "urn:uuid:")
	guid.encode(b[urnOffset:])
	return guid.str(b)
}

// Variant parses the variant from the UUID.
func (guid UUID) Variant() Variant {
	return Variant(guid[variantByte] & 0xE0) // byte & 11100000
}

// Version extracts the version from the UUID.
func (guid UUID) Version() Version {
	return Version(guid[versionByte] & 0xF0) // byte & 11110000
}

func (guid UUID) encode(dst []byte) {
	hex.Encode(dst, guid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], guid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], guid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], guid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], guid[10:])
}

func (guid UUID) str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
