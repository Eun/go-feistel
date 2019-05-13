package feistel

import (
	"bytes"
	"encoding/binary"
	"io"
)

type cbc struct{}

var CBC cbc

// EncryptReader reads data from an reader and writes the encrypted data to the writer
func (cbc) EncryptReader(r io.Reader, w io.Writer, rounds int, keys []uint32, previousBlock uint32) error {
	var block [8]byte
	prevLeft := uint32(previousBlock>>32) & 0xFFFFFFFF
	prevRight := uint32(previousBlock) & 0xFFFFFFFF
	for {
		left, right, err := readInt(r, block)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		left ^= prevLeft
		right ^= prevRight

		left, right = Encrypt(left, right, rounds, keys)
		if err := writeInt(w, block, left, right); err != nil {
			return err
		}
		prevLeft = left
		prevRight = right
	}
}

// Encrypt encrypts a provided buffer and returns it
func (cbc) Encrypt(buf []byte, rounds int, keys []uint32, previousBlock uint32) ([]byte, error) {
	var out bytes.Buffer
	if err := CBC.EncryptReader(bytes.NewBuffer(buf), &out, rounds, keys, previousBlock); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}


// EncryptUInt64 encrypts a provided uint64 and returns it
func (cbc) EncryptUInt64(n uint64, rounds int, keys []uint32, previousBlock uint32) (uint64, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	var err error
	buf, err = CBC.Encrypt(buf, rounds, keys, previousBlock)
	return binary.BigEndian.Uint64(buf), err
}

// EncryptInt64 encrypts a provided int64 and returns it
func (cbc) EncryptInt64(n int64, rounds int, keys []uint32, previousBlock uint32) (int64, error) {
	i, err := CBC.EncryptUInt64(uint64(n), rounds, keys, previousBlock)
	return (int64(i)), err
}

// DecryptReader reads data from an reader and writes the decrypted data to the writer
func (cbc) DecryptReader(r io.Reader, w io.Writer, rounds int, keys []uint32, previousBlock uint32) error {
	var block [8]byte
	var savedLeft uint32
	var savedRight uint32
	prevLeft := uint32(previousBlock>>32) & 0xFFFFFFFF
	prevRight := uint32(previousBlock) & 0xFFFFFFFF
	for {
		left, right, err := readInt(r, block)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		savedLeft = left
		savedRight = right
		left, right = Decrypt(left, right, rounds, keys)
		left ^= prevLeft
		right ^= prevRight
		if err := writeInt(w, block, left, right); err != nil {
			return err
		}
		prevLeft = savedLeft
		prevRight = savedRight
	}
}

// Decrypt decrypts a provided buffer and returns it
func (cbc) Decrypt(buf []byte, rounds int, keys []uint32, previousBlock uint32) ([]byte, error) {
	var out bytes.Buffer
	if err := CBC.DecryptReader(bytes.NewBuffer(buf), &out, rounds, keys, previousBlock); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

// DecryptUInt64 decrypts a provided uint64 and returns it
func (cbc) DecryptUInt64(n uint64, rounds int, keys []uint32, previousBlock uint32) (uint64, error) {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, n)
	var err error
	buf, err = CBC.Decrypt(buf, rounds, keys, previousBlock)
	return binary.BigEndian.Uint64(buf), err
}

// DecryptUnt64 decrypts a provided int64 and returns it
func (cbc) DecryptInt64(n int64, rounds int, keys []uint32, previousBlock uint32) (int64, error) {
	i, err := CBC.DecryptUInt64(uint64(n), rounds, keys, previousBlock)
	return (int64(i)), err
}
