package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Hex2Base64 takes a hexadecimal string and returns a the same value in
// base64 representation
func Hex2Base64(hexadecimal string) string {
	// convert the hexidecimal string to a byte array, and make sure that there are an even number of bytes
	if len(hexadecimal)%2 != 0 {
		hexadecimal = "0" + hexadecimal
	}
	hexBytes := []byte(hexadecimal)

	// convert the hexidecimal byte array to a decimal byte array
	bytes := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(bytes, hexBytes)
	if err != nil {
		fmt.Println(err)
	}

	// convert the decimal byte array into a base64 byte array
	b64Bytes := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
	base64.StdEncoding.Encode(b64Bytes, bytes)

	return string(b64Bytes)
}
