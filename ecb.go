package feistel

import (
	"bytes"
	"encoding/binary"
	"io"
)

type ecb struct{}

var ECB ecb

// EncryptReader reads data from an reader and writes the encrypted data to the writer
func (ecb) EncryptReader(r io.Reader, w io.Writer, rounds int, keys []uint32) error {
	var block [8]byte
	for {
		left, right, err := readInt(r, block)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		left, right = Encrypt(left, right, rounds, keys)
		if err := writeInt(w, block, left, right); err != nil {
			return err
		}
	}
}

// Encrypt encrypts a provided buffer and returns it
func (ecb) Encrypt(buf []byte, rounds int, keys []uint32) ([]byte, error) {
	var out bytes.Buffer
	if err := ECB.EncryptReader(bytes.NewBuffer(buf), &out, rounds, keys); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// EncryptUInt64 encrypts a provided uint64 and returns it
func (ecb) EncryptUInt64(n uint64, rounds int, keys []uint32) (uint64, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	var err error
	buf, err = ECB.Encrypt(buf, rounds, keys)
	return binary.BigEndian.Uint64(buf), err
}

// EncryptInt64 encrypts a provided int64 and returns it
func (ecb) EncryptInt64(n int64, rounds int, keys []uint32) (int64, error) {
	i, err := ECB.EncryptUInt64(uint64(n), rounds, keys)
	return (int64(i)), err
}

// DecryptReader reads data from an reader and writes the decrypted data to the writer
func (ecb) DecryptReader(r io.Reader, w io.Writer, rounds int, keys []uint32) error {
	var block [8]byte
	for {
		left, right, err := readInt(r, block)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		left, right = Decrypt(left, right, rounds, keys)
		if err := writeInt(w, block, left, right); err != nil {
			return err
		}
	}
}

// Decrypt decrypts a provided buffer and returns it
func (ecb) Decrypt(buf []byte, rounds int, keys []uint32) ([]byte, error) {
	var out bytes.Buffer
	if err := ECB.DecryptReader(bytes.NewBuffer(buf), &out, rounds, keys); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// DecryptUInt64 decrypts a provided uint64 and returns it
func (ecb) DecryptUInt64(n uint64, rounds int, keys []uint32) (uint64, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	var err error
	buf, err = ECB.Decrypt(buf, rounds, keys)
	return binary.BigEndian.Uint64(buf), err
}

// DecryptInt64 decrypts a provided int64 and returns it
func (ecb) DecryptInt64(n int64, rounds int, keys []uint32) (int64, error) {
	i, err := ECB.DecryptUInt64(uint64(n), rounds, keys)
	return (int64(i)), err
}
