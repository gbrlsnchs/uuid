package uuid_test

import (
	"math/rand"
	"regexp"
	"testing"

	. "github.com/gbrlsnchs/uuid"
)

func TestUUID(t *testing.T) {
	regex, err := regexp.Compile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		version Version
		factory func() (UUID, error)
	}{
		{
			Version1,
			func() (UUID, error) {
				return CreateV1()
			},
		},
		{
			Version2,
			func() (UUID, error) {
				id := uint32(rand.Intn(int(^uint32(0))))
				ldn := uint8(rand.Intn(int(^uint8(0))))
				return CreateV2(id, ldn)
			},
		},
		{
			Version3,
			func() (UUID, error) {
				guid, err := CreateV4()
				if err != nil {
					return Null, err
				}
				return CreateV3(guid, nil)
			},
		},
		{
			Version4,
			func() (UUID, error) {
				return CreateV4()
			},
		},
		{
			Version5,
			func() (UUID, error) {
				guid, err := CreateV4()
				if err != nil {
					return Null, err
				}
				return CreateV5(guid, nil)
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
				if want, got := true, regex.MatchString(guid.String()); want != got {
					t.Errorf("want %t, got %t", want, got)
				}
				t.Log(guid.String())
			}
		})
	}
}
