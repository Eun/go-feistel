package feistel

import (
	"bytes"
	"testing"
)

func TestCBCEncryptDecrypt(t *testing.T) {
	tests := []struct {
		Input         []byte
		Rounds        int
		Keys          []uint32
		PreviousBlock uint64
	}{
		{[]byte("Hello World"), 8, []uint32{0xDEADBEEF}, 0xFEEDFACE},
	}

	for _, test := range tests {
		enc, err := CBC.Encrypt(test.Input, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := CBC.Decrypt(enc, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if !bytes.Equal(test.Input, dec) {
			t.Fatalf("expected %v but got %v", test.Input, dec)
		}
	}
}

func TestCBCEncryptDecryptInt64(t *testing.T) {
	tests := []struct {
		Input         int64
		Rounds        int
		Keys          []uint32
		PreviousBlock uint64
	}{
		{123456789, 8, []uint32{0xDEADBEEF}, 0xFEEDFACE},
	}

	for _, test := range tests {
		enc, err := CBC.EncryptInt64(test.Input, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := CBC.DecryptInt64(enc, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if test.Input != dec {
			t.Fatalf("expected %d but got %d", test.Input, dec)
		}
	}
}

func TestCBCEncryptDecryptUInt64(t *testing.T) {
	tests := []struct {
		Input         uint64
		Rounds        int
		Keys          []uint32
		PreviousBlock uint64
	}{
		{123456789, 8, []uint32{0xDEADBEEF}, 0xFEEDFACE},
	}

	for _, test := range tests {
		enc, err := CBC.EncryptUInt64(test.Input, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := CBC.DecryptUInt64(enc, test.Rounds, test.Keys, test.PreviousBlock)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if test.Input != dec {
			t.Fatalf("expected %d but got %d", test.Input, dec)
		}
	}
}
