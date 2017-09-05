package main

import "testing"

func TestHexXOR(t *testing.T) {
	hex1 := "1c0111001f010100061a024b53535009181c"
	hex2 := "686974207468652062756c6c277320657965"
	XORResult := "746865206b696420646f6e277420706c6179"
	HexXORTestResult, err := HexXOR(hex1, hex2)
	if err != nil {
		t.Errorf("Hexidecimal XOR was incorrect, raised error: %s", err)
	}
	if HexXORTestResult != XORResult {
		t.Errorf("Hexidecimal XOR was incorrect, got: %s, want: %s", HexXORTestResult, XORResult)
	}
}
