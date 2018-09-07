package uuid

import (
	"crypto/md5"
)

// GenerateV3 generates a version 3 UUID based on a namespace UUID and additional data.
func GenerateV3(nspace UUID, data []byte) (UUID, error) {
	guid, err := hashUUID(md5.New(), nspace, data)
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version3), nil
}
