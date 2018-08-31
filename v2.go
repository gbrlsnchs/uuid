package uuid

import "encoding/binary"

// CreateV2 generates a version 2 UUID.
//
// It basically returns a version 1 UUID but overrides
// the least significant 8 bits of the clock sequence by
// the variable "ldn", a local domain name, and the least significant
// 32 bits of the timestamp an integer identifier "id".
func CreateV2(id uint32, ldn uint8) (UUID, error) {
	guid, err := timestampUUID()
	if err != nil {
		return Null, err
	}
	binary.BigEndian.PutUint32(guid[:4], id)
	guid[9] = ldn
	guid.setVersion(Version2)
	return guid, nil
}
