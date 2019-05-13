package feistel

import (
	"testing"

	"bytes"
)

func Test_findEnd8(t *testing.T) {
	tests := []struct {
		name string
		args [8]byte
		want []byte
	}{
		{"No input", [8]byte{}, []byte{}},
		{"No End", [8]byte{0, 1, 2, 3, 4, 5, 6, 7}, []byte{0, 1, 2, 3, 4, 5, 6, 7}},
		{"End after 3", [8]byte{0, 1, 2, 3, 0, 0, 0}, []byte{0, 1, 2, 3}},
		{"Alternate Zero/One", [8]byte{0, 1, 0, 1, 0, 1, 0}, []byte{0, 1, 0, 1, 0, 1}},
		{"All Zero", [8]byte{0, 0, 0, 0, 0, 0, 0}, []byte{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := findEnd8(tt.args)
			if !bytes.Equal(tt.want[:], tt.args[:n]) {
				t.Errorf("findEnd() = %v, want %v", tt.args[:n], tt.want)
			}
		})
	}
}
