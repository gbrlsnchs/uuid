package uuid

import "encoding/binary"

// GenerateV2 generates a version 2 UUID.
//
// It basically returns a version 1 UUID but overrides
// the least significant 8 bits of the clock sequence by
// the variable "ldn", a local domain name, and the least significant
// 32 bits of the timestamp an integer identifier "id".
func GenerateV2(id uint32, ldn uint8, random bool) (UUID, error) {
	guid, err := timestampUUID(random)
	if err != nil {
		return Null, err
	}
	binary.BigEndian.PutUint32(guid[:4], id)
	guid[9] = ldn
	return guid.withVersion(Version2), nil
}
