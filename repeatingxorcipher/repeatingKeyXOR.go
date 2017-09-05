package repeatingxorcipher

import "encoding/hex"

// RepeatingKeyXOREncrypt takes a plaintext message and a repeating key and applies the repeating
// key to the plaintext message
// @param plaintext: a string representing the message to encrypt
// @param key: an array of bytes that is the repeating key
// @return: a hexadecimal string representing the encrypted message
func RepeatingKeyXOREncrypt(plaintext string, key []byte) string {
	unencryptedBytes := []byte(plaintext)
	encryptedMessage := RepeatingKeyXOR(unencryptedBytes, key)
	return hex.EncodeToString(encryptedMessage)
}

// RepeatingKeyXOR takes a byte array message and applies a repeating key to it
// this can be used to encrypt and decrypt in repeating key xor cipher
// @param message: a byte array representing the original message
// @param key: a byte array representing the rotating key
// @return: a byte array that is the result of applying the rotating key to the message
func RepeatingKeyXOR(message, key []byte) []byte {
	var repeatingKeyIndex int
	resultingMessage := make([]byte, len(message))

	for i := 0; i < len(message); i++ {
		resultingMessage[i] = message[i] ^ key[repeatingKeyIndex]

		if repeatingKeyIndex+1 >= len(key) {
			repeatingKeyIndex = 0
		} else {
			repeatingKeyIndex++
		}
	}

	return resultingMessage
}
