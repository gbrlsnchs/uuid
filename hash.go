package uuid

import "hash"

var (
	// NamespaceDNS is the namespace UUID defined for DNS.
	NamespaceDNS, _ = Parse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	// NamespaceURL is the namespace UUID defined for URL.
	NamespaceURL, _ = Parse("6ba7b811-9dad-11d1-80b4-00c04fd430c8")
	// NamespaceOID is the namespace UUID defined for ISO OID.
	NamespaceOID, _ = Parse("6ba7b812-9dad-11d1-80b4-00c04fd430c8")
	// NamespaceX500 is the namespace UUID defined for X.500 DN.
	NamespaceX500, _ = Parse("6ba7b814-9dad-11d1-80b4-00c04fd430c8")
)

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
