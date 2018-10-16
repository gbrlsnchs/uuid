package uuid

import "crypto/md5"

// GenerateV3 generates a version 3 UUID based on a namespace UUID and additional data.
func GenerateV3(ns UUID, data []byte) (UUID, error) {
	guid, err := hashUUID(md5.New(), ns, data)
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version3), nil
}

// V3 returns a version 3 UUID or panics otherwise.
func V3(ns UUID, data []byte) UUID {
	return uuidOrPanic(GenerateV3(ns, data))
}
