package feistel

import (
	"bytes"
	"testing"
)

func TestECBEncryptDecrypt(t *testing.T) {
	tests := []struct {
		Input  []byte
		Rounds int
		Keys   []uint32
	}{
		{[]byte("Hello World"), 8, []uint32{0xDEADBEEF}},
	}

	for _, test := range tests {
		enc, err := ECB.Encrypt(test.Input, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := ECB.Decrypt(enc, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if !bytes.Equal(test.Input, dec) {
			t.Fatalf("expected %v but got %v", test.Input, dec)
		}
	}
}

func TestECBEncryptDecryptInt64(t *testing.T) {
	tests := []struct {
		Input  int64
		Rounds int
		Keys   []uint32
	}{
		{123456789, 8, []uint32{0xDEADBEEF}},
	}

	for _, test := range tests {
		enc, err := ECB.EncryptInt64(test.Input, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := ECB.DecryptInt64(enc, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if test.Input != dec {
			t.Fatalf("expected %d but got %d", test.Input, dec)
		}
	}
}

func TestECBEncryptDecryptUInt64(t *testing.T) {
	tests := []struct {
		Input  uint64
		Rounds int
		Keys   []uint32
	}{
		{123456789, 8, []uint32{0xDEADBEEF}},
	}

	for _, test := range tests {
		enc, err := ECB.EncryptUInt64(test.Input, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		dec, err := ECB.DecryptUInt64(enc, test.Rounds, test.Keys)
		if err != nil {
			t.Fatalf("got error: %v", err)
		}
		if test.Input != dec {
			t.Fatalf("expected %d but got %d", test.Input, dec)
		}
	}
}
