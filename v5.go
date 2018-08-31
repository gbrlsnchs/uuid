package uuid

import (
	"crypto/sha1"
)

// CreateV5 generates a version 5 UUID based on a namespace UUID and additional data.
func CreateV5(nspace UUID, data []byte) (UUID, error) {
	return hashUUID(sha1.New(), nspace, data, Version5)
}
