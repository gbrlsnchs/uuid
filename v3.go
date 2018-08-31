package uuid

import (
	"crypto/md5"
)

// CreateV3 generates a version 3 UUID based on a namespace UUID and additional data.
func CreateV3(nspace UUID, data []byte) (UUID, error) {
	return hashUUID(md5.New(), nspace, data, Version3)
}
