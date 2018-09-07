package uuid

import (
	"crypto/rand"
	"io"
)

// GenerateV4 generates a version 4 UUID.
func GenerateV4() (UUID, error) {
	var guid UUID
	_, err := io.ReadFull(rand.Reader, guid[:])
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version4), nil
}
