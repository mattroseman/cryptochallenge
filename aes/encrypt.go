package aes

import "errors"

// Encrypt takes some plaintext and uses AES encryption to encrypt the stirng,
// and return the encrypted string
func Encrypt(plaintext string, key []byte) (string, error) {
	if len(key) != 16 {
		return "", errors.New("Keysize must be 128 bits/16 bytes")
	}
}

// encryptBlock takes a 128 bit block of plaintext and encrypts it using the 128 bit key
func encryptBlock(block []byte, key []byte) ([]byte, error) {
	if len(key) != 16 {
		return "", errors.New("Keysize must be 128 bits/16 bytes")
	}
	if len(block) != 16 {
		return "", errors.New("Blocksize must be 128 bits/16 bytes")
	}

	xorBlock(block, key)
}

func xorBlock(a, b []byte) error {
	if len(a) != len(b) {
		return errors.New("Blocks must be the same length")
	}

	for i := range a {
		a[i] ^= b[i]
	}

	return nil
}
