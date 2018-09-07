package uuid

import (
	"crypto/sha1"
)

// GenerateV5 generates a version 5 UUID based on a namespace UUID and additional data.
func GenerateV5(nspace UUID, data []byte) (UUID, error) {
	guid, err := hashUUID(sha1.New(), nspace, data)
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version5), nil
}
