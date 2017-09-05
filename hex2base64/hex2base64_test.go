package main

import "testing"

func TestHex2Base64(t *testing.T) {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	converted := Hex2Base64(hexString)
	if converted != base64String {
		t.Errorf("Hexidecimal to Base64 conversion was incorrect, got: %s, want: %s", converted, base64String)
	}
}
