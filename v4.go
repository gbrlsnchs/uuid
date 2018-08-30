package uuid

import (
	"crypto/rand"
	"io"
)

// V4 generates a version 4 UUID.
func V4() (UUID, error) {
	var guid UUID
	_, err := io.ReadFull(rand.Reader, guid[:])
	if err != nil {
		return Null, err
	}
	guid.setVersion(Version4)
	return guid, nil
}
