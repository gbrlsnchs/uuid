package uuid

// Variant is a UUID variant as per the RFC 4122, § 4.1.1.
type Variant byte

const (
	// VariantNCS is a reserved variant for NCS backward compatibility.
	VariantNCS Variant = 0x00
	// VariantRFC is the current variant defined by the RFC 4122.
	VariantRFC Variant = 0x80 // 10xxxxxx
	// VariantMicrosoft is a reserved variant for Microsoft Corporation backward compatibility.
	VariantMicrosoft Variant = 0xC0 // 110xxxxx
	// VariantUndefined is a reserved variant for future definition (still undefined).
	VariantUndefined Variant = 0xE0 // 111xxxxx
	variantByte              = 8
)

func (v Variant) String() string {
	switch {
	case v&VariantRFC > 0:
		return "RFC 4122"
	case v&VariantMicrosoft > 0:
		return "Microsoft"
	case v&VariantUndefined > 0:
		return "Undefined"
	default:
		return "NCS"
	}
}
