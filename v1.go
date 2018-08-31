package uuid

// CreateV1 generates a version 1 UUID.
func CreateV1(random bool) (UUID, error) {
	guid, err := timestampUUID(random)
	if err != nil {
		return Null, err
	}
	guid.setVersion(Version1)
	return guid, nil
}
