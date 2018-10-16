package uuid_test

import (
	"crypto/rand"
	prng "math/rand"
	"regexp"
	"testing"
	"time"

	. "github.com/gbrlsnchs/uuid"
)

func TestUUID(t *testing.T) {
	prng.Seed(time.Now().Unix())
	regexpMap := map[Version]*regexp.Regexp{
		Version1: regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-1[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"),
		Version2: regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-2[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"),
		Version3: regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-3[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"),
		Version4: regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"),
		Version5: regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-5[0-9a-fA-F]{3}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$"),
	}
	testCases := []struct {
		version Version
		factory func() (UUID, error)
	}{
		{
			Version1,
			func() (UUID, error) {
				return GenerateV1(false)
			},
		},
		{
			Version1,
			func() (UUID, error) {
				return GenerateV1(true)
			},
		},
		{
			Version2,
			func() (UUID, error) {
				id := uint32(prng.Intn(int(^uint32(0))))
				ldn := uint8(prng.Intn(int(^uint8(0))))
				return GenerateV2(id, ldn, false)
			},
		},
		{
			Version2,
			func() (UUID, error) {
				id := uint32(prng.Intn(int(^uint32(0))))
				ldn := uint8(prng.Intn(int(^uint8(0))))
				return GenerateV2(id, ldn, true)
			},
		},
		{
			Version3,
			func() (UUID, error) {
				guid, err := GenerateV4(rand.Reader)
				if err != nil {
					return Null, err
				}
				return GenerateV3(guid, nil)
			},
		},
		{
			Version4,
			func() (UUID, error) {
				return GenerateV4(rand.Reader)
			},
		},
		{
			Version5,
			func() (UUID, error) {
				guid, err := GenerateV4(rand.Reader)
				if err != nil {
					return Null, err
				}
				return GenerateV5(guid, nil)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.version.String(), func(t *testing.T) {
			for i := 0; i < 1000; i++ {
				guid, err := tc.factory()
				if err != nil {
					t.Error(err)
				}
				if want, got := tc.version, guid.Version(); want != got {
					t.Errorf("want %s, got %s", want.String(), got.String())
				}
				if want, got := VariantRFC4122, guid.Variant()&0xC0; want != got { // filter only the MS 2 bits of variant
					t.Errorf("want %s, got %s", want.String(), got.String())
				}
				if want, got := true, regexpMap[guid.Version()].MatchString(guid.String()); want != got {
					t.Errorf("want %t, got %t", want, got)
				}
			}
		})
	}
}

func TestUUIDScan(t *testing.T) {
	testCases := []struct {
		guid UUID
	}{
		{Null},
		{NamespaceDNS},
		{NamespaceURL},
		{NamespaceOID},
		{NamespaceX500},
	}
	for _, tc := range testCases {
		t.Run(tc.guid.String(), func(t *testing.T) {
			// Test it as a 16-byte array.
			var guid UUID
			if want, got := (error)(nil), guid.Scan(tc.guid.Bytes()); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid; want != got {
				t.Errorf("want %v, got %v", want, got)
			}

			// Test it as a 36-byte array encoded to hex.
			s := tc.guid.String()
			guid = Null
			if want, got := (error)(nil), guid.Scan([]byte(s)); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid; want != got {
				t.Errorf("want %v, got %v", want, got)
			}

			// Test it as a 36-byte string encoded to hex.
			guid = Null
			if want, got := (error)(nil), guid.Scan(s); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid; want != got {
				t.Errorf("want %v, got %v", want, got)
			}
		})
	}
}
