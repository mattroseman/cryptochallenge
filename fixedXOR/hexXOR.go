package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// HexXOR takes two hexadecimal numbers and XOR's them and returns the result
func HexXOR(firstHex, secondHex string) (string, error) {
	// convert the hex strings to bytes
	hex1 := []byte(firstHex)
	hex2 := []byte(secondHex)

	// check that the two hex values are of the same length
	if len(hex1) != len(hex2) {
		return "", errors.New("The lengths of the hex numbers must be equal")
	}

	// convert the hex values to decimal values
	bytes1 := make([]byte, hex.DecodedLen(len(hex1)))
	bytes2 := make([]byte, hex.DecodedLen(len(hex2)))
	_, err := hex.Decode(bytes1, hex1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = hex.Decode(bytes2, hex2)
	if err != nil {
		fmt.Println(err)
	}

	// XOR the two decimal byte values
	XORValue := make([]byte, len(bytes1))
	for i := range bytes1 {
		XORValue[i] = bytes1[i] ^ bytes2[i]
	}

	// convert the decimal XOR value back to hexidecimal
	XORHexValue := make([]byte, hex.EncodedLen(len(XORValue)))
	_ = hex.Encode(XORHexValue, XORValue)

	return string(XORHexValue), nil
}
