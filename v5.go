package uuid

import "crypto/sha1"

// GenerateV5 generates a version 5 UUID based on a namespace UUID and additional data.
func GenerateV5(ns UUID, data []byte) (UUID, error) {
	guid, err := hashUUID(sha1.New(), ns, data)
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version5), nil
}

// V5 returns a version 5 UUID or panics otherwise.
func V5(ns UUID, data []byte) UUID {
	return uuidOrPanic(GenerateV5(ns, data))
}
