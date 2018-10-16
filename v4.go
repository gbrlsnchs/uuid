package uuid

import "io"

// GenerateV4 generates a version 4 UUID.
func GenerateV4(r io.Reader) (UUID, error) {
	var guid UUID
	_, err := io.ReadFull(r, guid[:])
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version4), nil
}

// V4 returns a version 4 UUID or panics otherwise.
func V4(r io.Reader) UUID {
	return uuidOrPanic(GenerateV4(r))
}
