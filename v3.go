package uuid

import (
	"crypto/md5"
)

// V3 generates a version 3 UUID based on a namespace UUID and additional data.
func V3(nspace UUID, data []byte) (UUID, error) {
	return hashUUID(md5.New(), nspace, data, Version3)
}
