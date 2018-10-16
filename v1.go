package uuid

// GenerateV1 generates a version 1 UUID.
func GenerateV1(random bool) (UUID, error) {
	guid, err := timestampUUID(random)
	if err != nil {
		return Null, err
	}
	return guid.withVersion(Version1), nil
}

// V1 returns a version 1 UUID or panics otherwise.
func V1(random bool) UUID {
	return uuidOrPanic(GenerateV1(random))
}
