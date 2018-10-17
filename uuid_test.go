package uuid_test

import (
	"bytes"
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

func TestUUIDGUID(t *testing.T) {
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
		t.Run(tc.guid.GUID(), func(t *testing.T) {
			if want, got := "{"+tc.guid.String()+"}", tc.guid.GUID(); want != got {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestUUIDMarshal(t *testing.T) {
	testCases := []struct {
		guid           UUID
		expectedBinary []byte
		expectedJSON   string
		expectedText   string
	}{
		{Null, nil, "", ""},
		{NamespaceDNS, NamespaceDNS[:], NamespaceDNS.String(), NamespaceDNS.String()},
		{NamespaceURL, NamespaceURL[:], NamespaceURL.String(), NamespaceURL.String()},
		{NamespaceOID, NamespaceOID[:], NamespaceOID.String(), NamespaceOID.String()},
		{NamespaceX500, NamespaceX500[:], NamespaceX500.String(), NamespaceX500.String()},
	}

	for _, tc := range testCases {
		t.Run(tc.guid.String(), func(t *testing.T) {
			var (
				b   []byte
				err error
			)

			b, err = tc.guid.MarshalBinary()
			if want, got := (error)(nil), err; want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.expectedBinary, b; !bytes.Equal(want, got) {
				t.Errorf("want %b, got %b", want, got)
			}

			b, err = tc.guid.MarshalJSON()
			if want, got := (error)(nil), err; want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.expectedJSON, b; want != string(got) {
				t.Errorf("want %s, got %s", want, got)
			}

			b, err = tc.guid.MarshalText()
			if want, got := (error)(nil), err; want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.expectedText, b; want != string(got) {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}

func TestUUIDUnmarshal(t *testing.T) {
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
			var guid2 UUID
			if want, got := (error)(nil), guid2.UnmarshalBinary(tc.guid[:]); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid2; want != got {
				t.Errorf("want %v, got %v", want, got)
			}

			b := []byte(tc.guid.String())
			guid2 = Null
			if want, got := (error)(nil), guid2.UnmarshalJSON(b); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid2; want != got {
				t.Errorf("want %v, got %v", want, got)
			}

			guid2 = Null
			if want, got := (error)(nil), guid2.UnmarshalText(b); want != got {
				t.Errorf("want %v, got %v", want, got)
			}
			if want, got := tc.guid, guid2; want != got {
				t.Errorf("want %v, got %v", want, got)
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
			if want, got := (error)(nil), guid.Scan(tc.guid[:]); want != got {
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
