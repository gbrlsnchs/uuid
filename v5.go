package uuid

import (
	"crypto/sha1"
)

func generateV5(nspace UUID, data []byte) (UUID, error) {
	return hashUUID(sha1.New(), nspace, data)
}
