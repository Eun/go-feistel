package feistel

import "testing"

func BenchmarkMemClr8(b *testing.B) {
	b.Run("for range", func(b *testing.B) {
		clr := func(b [8]byte) {
			for i := range b {
				b[i] = zero
			}
		}

		var buf [8]byte
		for i := 0; i < b.N; i++ {
			clr(buf)
		}
	})

	b.Run("for i loop", func(b *testing.B) {
		clr := func(b [8]byte) {
			for i := 0; i < 8; i++ {
				b[i] = zero
			}
		}

		var buf [8]byte
		for i := 0; i < b.N; i++ {
			clr(buf)
		}
	})
}
