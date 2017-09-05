package xorcipher

import (
	"encoding/hex"
	"fmt"
)

// UnencryptXOREncryptedHexString takes an encrypted string in hexadecimal form, that was encrypted with a single XOR
// byte, and finds the most likely byte use assuming the original plaintext message was in english
// @param encryptedString: a hexadecimal string
// @param return: the best guess at what the original plaintext message is
func UnencryptXOREncryptedHexString(encryptedString string) string {
	// make sure the encrypted string has an even number of hexidecimal values
	if len(encryptedString)%2 == 1 {
		encryptedString = "0" + encryptedString
	}

	// convert the hexidecimal string to a byte array
	encryptedHex := []byte(encryptedString)

	// convert the hexidecimal byte array to a decimal byte array
	encryptedBytes := make([]byte, hex.DecodedLen(len(encryptedHex)))
	_, err := hex.Decode(encryptedBytes, encryptedHex)
	if err != nil {
		fmt.Println(err)
	}

	_, unencryptedString := FindXORByte(encryptedBytes)

	return unencryptedString
}
