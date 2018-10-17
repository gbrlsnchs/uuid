package uuid

import (
	"crypto/rand"
	"io"
)

// GenerateV4 generates a version 4 UUID.
//
// If a nil reader is passed as argument, crypto/rand.Reader is used instead.
func GenerateV4(r io.Reader) (UUID, error) {
	if r == nil {
		r = rand.Reader
	}
	var guid UUID
	_, err := io.ReadFull(r, guid[:])
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version4), nil
}

// V4 returns a version 4 UUID or panics otherwise.
//
// If a nil reader is passed as argument, crypto/rand.Reader is used instead.
func V4(r io.Reader) UUID {
	return uuidOrPanic(GenerateV4(r))
}
