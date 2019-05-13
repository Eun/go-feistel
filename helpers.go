package feistel

import (
	"fmt"
	"io"
)

const zero = 0

// special functions to operate with [8]byte
// memclr8 clears a byte slice
func memclr8(b [8]byte) {
	for i := range b {
		b[i] = zero
	}
}

// findEnd8 finds the last zeros in an byte slice
func findEnd8(b [8]byte) int {
	for i := 7; i >= 0; i-- {
		if b[i] != zero {
			return i + 1
		}
	}
	return 0
}

// readInt reads the next left and right part from an reader
// pass the reader and the buffer that should be used to read the int
func readInt(r io.Reader, buf [8]byte) (left, right uint32, err error) {
	memclr8(buf)
	n, err := r.Read(buf[:])
	if err != nil {
		return 0, 0, err
	}
	if n == 0 {
		return 0, 0, io.ErrUnexpectedEOF
	}

	return uint32(buf[4]) | uint32(buf[5])<<8 | uint32(buf[6])<<16 | uint32(buf[7])<<24, uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24, nil
}

// writeInt writes an integer to an writer
// pass the writer and the buffer that should be used for writing
func writeInt(w io.Writer, buf [8]byte, left, right uint32) error {
	buf[0] = byte(right)
	buf[1] = byte(right >> 8)
	buf[2] = byte(right >> 16)
	buf[3] = byte(right >> 24)
	buf[4] = byte(left)
	buf[5] = byte(left >> 8)
	buf[6] = byte(left >> 16)
	buf[7] = byte(left >> 24)

	end := findEnd8(buf)
	n, err := w.Write(buf[:end])
	if err != nil {
		return err
	}
	if end != n {
		return fmt.Errorf("expected to write %d got %d bytes", end, n)
	}
	return nil
}
