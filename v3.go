package uuid

import (
	"crypto/md5"
)

func generateV3(nspace UUID, data []byte) (UUID, error) {
	return hashUUID(md5.New(), nspace, data)
}
