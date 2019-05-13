package feistel

func Encrypt(left, right uint32, rounds int, keys []uint32) (uint32, uint32) {
	size := len(keys)
	for i := 0; i < rounds; i++ {
		left1 := (left ^ keys[i%size]) ^ right
		right1 := left
		if i == (rounds - 1) {
			left = right1
			right = left1
		} else {
			left = left1
			right = right1
		}
	}
	return left, right
}

func Decrypt(left, right uint32, rounds int, keys []uint32) (uint32, uint32) {
	size := len(keys)
	for i := 0; i < rounds; i++ {
		left1 := (left ^ keys[(rounds-i-1)%size]) ^ right
		right1 := left
		if i == (rounds - 1) {
			left = right1
			right = left1
		} else {
			left = left1
			right = right1
		}
	}
	return left, right
}
