package uuid

import (
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
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
	// ErrInvalid is the error for invalid UUID formats.
	ErrInvalid = errors.New("uuid: invalid uuid")
)

// UUID is a 16-byte array Universally Unique IDentifier, as per the RFC 4122.
type UUID [byteSize]byte

// Parse parses a UUID 36-byte slice encoded in hexadecimal and converts it to a 16-byte array.
func Parse(s string) (UUID, error) {
	if len(s) != hexSize {
		if len(s) != urnSize {
			return Null, ErrInvalid
		}
		if !strings.HasPrefix(s, urnPrefix) {
			return Null, ErrInvalid
		}
		s = s[urnOffset:]
	}
	var (
		b    = []byte(s)
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

func uuidOrPanic(guid UUID, err error) UUID {
	if err != nil {
		panic(err)
	}
	return guid
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
	b := microsoft[:]
	return unsafeStr(&b)
}

// IsNull returns whether the UUID is a null UUID.
func (guid UUID) IsNull() bool {
	return guid == Null
}

// MarshalBinary implements binary marshaling.
func (guid UUID) MarshalBinary() ([]byte, error) {
	return guid.Bytes(), nil // 16-byte array
}

// MarshalJSON implements JSON marshaling.
func (guid UUID) MarshalJSON() ([]byte, error) {
	return json.Marshal(guid.String()) // 36-byte hex-encoded string
}

// MarshalText implements text marshaling.
func (guid UUID) MarshalText() ([]byte, error) {
	var xid [hexSize]byte
	b := xid[:]
	guid.encode(b)
	return b, nil // 36-byte hex-encoded slice
}

// Scan implements scanning a UUID from an SQL database.
func (guid *UUID) Scan(v interface{}) error {
	switch vv := v.(type) {
	case []byte:
		if len(vv) == byteSize { // byte array in binary form
			copy(guid[:], vv)
			break
		}
		return guid.UnmarshalText(vv)
	case string:
		return guid.UnmarshalText([]byte(vv))
	default:
		return fmt.Errorf("uuid: cannot scan type %T into UUID", v)
	}
	return nil
}

// String converts the 16-byte UUID to a 36-byte string encoded in hexadecimal.
func (guid UUID) String() string {
	b, _ := guid.MarshalText()
	return unsafeStr(&b)
}

// UnmarshalBinary implements binary unmarshaling.
func (guid *UUID) UnmarshalBinary(b []byte) error {
	if len(b) != byteSize {
		return ErrInvalid
	}
	copy(guid.Bytes(), b)
	return nil
}

// UnmarshalJSON implements JSON unmarshaling.
func (guid *UUID) UnmarshalJSON(b []byte) error {
	return guid.UnmarshalText(b)
}

// UnmarshalText implements text unmarshaling.
func (guid *UUID) UnmarshalText(b []byte) error {
	u, err := Parse(unsafeStr(&b))
	if err != nil {
		return err
	}
	*guid = u
	return nil
}

// URN returns the UUID as string conformed to RFC 2141.
func (guid UUID) URN() string {
	var urn [urnSize]byte
	b := urn[:]
	copy(b, "urn:uuid:")
	guid.encode(b[urnOffset:])
	return unsafeStr(&b)
}

// Value implements saving the UUID as a 36-byte string encoded to hex to an SQL database.
func (guid UUID) Value() (driver.Value, error) {
	return driver.String.ConvertValue(guid.String())
}

// Variant parses the variant from the UUID.
func (guid UUID) Variant() Variant {
	v := Variant(guid[variantByte] & 0xE0) // byte & 11100000
	switch {
	case v&VariantRFC4122 > 0:
		return VariantRFC4122
	case v&VariantMicrosoft > 0:
		return VariantMicrosoft
	case v&VariantUndefined > 0:
		return VariantUndefined
	default:
		return VariantNCS
	}
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

func (guid UUID) withVersion(v Version) UUID {
	// Clear the 4 most significant bits and set version.
	guid[versionByte] = guid[versionByte]&0x0F | byte(v) // (byte & 00001111) | 01000000
	// Clear the 2 most significant bits and set variant.
	// As the RFC 4122 only takes into account 2 bits, the third most significant bit is ignored and thus not zeroed.
	guid[variantByte] = guid[variantByte]&0x3F | byte(VariantRFC4122) // (byte & 00111111) | 10000000
	return guid
}
