package feistel

// Encrypt takes a left and right uint32 and encrypts it using the Feistel cipher,
// it returns the encrypted left and right parts
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

// Decrypt takes a left and right uint32 and decrypts it using the Feistel cipher,
// it returns the decrypted left and right parts
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
