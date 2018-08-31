package uuid

// CreateV1 generates a version 1 UUID.
func CreateV1() (UUID, error) {
	guid, err := timestampUUID()
	if err != nil {
		return Null, err
	}
	guid.setVersion(Version1)
	return guid, nil
}
