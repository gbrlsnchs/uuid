package uuid

import (
	"crypto/rand"
	"io"
)

func generateV4() (UUID, error) {
	var guid UUID
	_, err := io.ReadFull(rand.Reader, guid[:])
	if err != nil {
		return Null, err
	}
	return guid, nil
}
