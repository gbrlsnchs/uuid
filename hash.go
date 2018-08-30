package uuid

import "hash"

func hashUUID(h hash.Hash, nspace UUID, data []byte, v Version) (UUID, error) {
	if _, err := h.Write(nspace[:]); err != nil {
		return Null, err
	}
	if _, err := h.Write(data); err != nil {
		return Null, err
	}
	sum := h.Sum(nil)
	var guid UUID
	copy(guid[:], sum)
	guid.setVersion(v)
	return guid, nil
}
