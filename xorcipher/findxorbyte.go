package xorcipher

import (
	"math"

	"github.com/mroseman95/cryptochallenge/letterdistribution"
)

// FindXORByte takes an enrypted byte array, and tries to guess what the byte used to XOR was.
// It assumes the decoded string is some english text.
// @param encryptedBytes: a byte array that represents the encrypted message
// @return: the best guess at what byte was used to encrypt, and the string representing the original
// plaintext message
func FindXORByte(encryptedBytes []byte) (byte, string) {
	// the byte that will XOR the encrypted message
	var xorByte byte
	encryptedLen := len(encryptedBytes)
	// the result of the encrypted message XOR'd with the xorByte
	xorCandidate := make([]byte, encryptedLen)
	// the diff with normal english character probability that is the current minimum
	minCandidateDiff := math.Inf(1)
	// the xorByte that results in the minimum diff in character probability with normal english text
	minDiffXORByte := xorByte
	// the string that was the result of encryptedBytes XOR minDiffXORByte
	var minDiffDecodedString string

	// try every possible xorByte
	for i := 0; i < 256; i++ {
		// XOR encryptedBytes with this xorByte
		for j := 0; j < encryptedLen; j++ {
			xorCandidate[j] = encryptedBytes[j] ^ xorByte
		}

		// convert this byte array to string
		xorStringCandidate := string(xorCandidate)
		candidateDiff := letterdistribution.GetEnglishCharacterProbability(xorStringCandidate)

		if candidateDiff < minCandidateDiff {
			minDiffXORByte = xorByte
			minCandidateDiff = candidateDiff
			minDiffDecodedString = xorStringCandidate
		}

		xorByte++
	}

	return minDiffXORByte, minDiffDecodedString
}
