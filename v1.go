package uuid

// V1 generates a version 1 UUID.
func V1() (UUID, error) {
	guid, err := timestampUUID()
	if err != nil {
		return Null, err
	}
	guid.setVersion(Version1)
	return guid, nil
}
