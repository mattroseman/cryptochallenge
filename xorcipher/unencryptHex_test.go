package xorcipher

import (
	"testing"
)

func TestFindXORByte(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	result := "Cooking MC's like a pound of bacon"
	decodedString := UnencryptXOREncryptedHexString(hexString)
	if decodedString != result {
		t.Errorf("Decoding the hex string found: %s, but wanted: %s", decodedString, result)
	}
}
