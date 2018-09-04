package uuid_test

import (
	"testing"

	. "github.com/gbrlsnchs/uuid"
)

func BenchmarkV4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GenerateV4(); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkVersion(b *testing.B) {
	guid, err := GenerateV4()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = guid.Version()
	}
}

func BenchmarkVariant(b *testing.B) {
	guid, err := GenerateV4()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = guid.Variant()
	}
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := Parse("d9ab3f01-482f-425d-8a10-a24b0abfe661"); err != nil {
			b.Error(err)
		}
	}
}
