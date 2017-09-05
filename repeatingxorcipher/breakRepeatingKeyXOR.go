package repeatingxorcipher

import (
	"encoding/base64"
	"math"

	"github.com/mroseman95/cryptochallenge/hamming"
	"github.com/mroseman95/cryptochallenge/letterdistribution"
	"github.com/mroseman95/cryptochallenge/xorcipher"
)

var maxKeySize = 40
var numKeySizeBlocks = 4
var numProbableKeySizes = 3

// BreakRepeatingKeyXOR takes an encrypted message represented in base64 that was encrypted using
// a repeating key XOR cypher on an english plaintext, and attempts to find the original plaintext message
// @param encryptedMessage: a string representing the encyrpted message
// @return: the english plaintext that is most likely to be the original message and an error value
func BreakRepeatingKeyXOR(encryptedMessage string) (string, error) {
	// convert base64 encoding into decoded bytes
	b64EncryptedBytes := []byte(encryptedMessage)
	encryptedBytes := make([]byte, base64.StdEncoding.DecodedLen(len(b64EncryptedBytes)))
	_, err := base64.StdEncoding.Decode(encryptedBytes, b64EncryptedBytes)
	if err != nil {
		return "", err
	}

	probableKeySizes := ProbableKeySizes(encryptedBytes)

	// minDecryptedMessageDist is smallest distance an candidate decrypted message has form english text
	minDecryptedMessageDist := math.Inf(1)
	// minDecryptedMessage is the cooresponding plaintext message that has the smallest dist from english text
	var minDecryptedMessage string

	for _, keySize := range probableKeySizes {
		candidateRotatingCipher := make([]byte, keySize)

		for i := 0; i < keySize; i++ {
			// take every keySize byte (starting at i) and solve for a single character XOR
			singleXOREncryptedBytes := make([]byte, int(math.Floor(float64(len(encryptedBytes)-i)/float64(keySize))))
			for j := range singleXOREncryptedBytes {
				singleXOREncryptedBytes[j] = encryptedBytes[i+j*keySize]
			}

			xorByte, _ := xorcipher.FindXORByte(singleXOREncryptedBytes)
			candidateRotatingCipher[i] = xorByte
		}

		candidateDecryptedMessage := string(RepeatingKeyXOR(encryptedBytes, candidateRotatingCipher))
		dist := letterdistribution.GetEnglishCharacterProbability(candidateDecryptedMessage)

		if dist < minDecryptedMessageDist {
			minDecryptedMessageDist = dist
			minDecryptedMessage = candidateDecryptedMessage
		}
	}

	return minDecryptedMessage, nil
}

// ProbableKeySizes takes some encrypted bytes, and determines what the probable key sizes used to encrypt
// that message.
// @param encryptedBytes: a byte array that contains the encrypted message
// @return: an array of integers that are the probable key sizes
func ProbableKeySizes(encryptedBytes []byte) []int {
	minNormalizedDists := make([]float64, numProbableKeySizes)
	minNormalizedDistKeySizes := make([]int, numProbableKeySizes)
	for i := range minNormalizedDists {
		minNormalizedDists[i] = math.Inf(1)
		minNormalizedDistKeySizes[i] = 0
	}

	for keySize := 2; keySize <= maxKeySize; keySize++ {
		// if 2 keySize blocks can't fit in the encrypted message, break out
		// and don't attempt this keysize
		if 2*keySize > len(encryptedBytes) {
			break
		}

		// Try and find the average difference between numKeySizeBlocks, but if they don't all fit, we know
		// at least two will fit. Just use the average of what will fit
		previousKeySizeBlock := encryptedBytes[:keySize]
		var totalKeySizeBlockDist int
		var numKeySizeBlocksUsed int

		// work through each block, until numKeySizeBlocks have been visited, or another keySize block won't fit in the encryptedBytes
		for i := keySize; i < int(math.Min(float64(keySize*numKeySizeBlocks), float64(len(encryptedBytes)-keySize))); i += keySize {
			keySizeBlock := encryptedBytes[i : i+keySize]

			// get distance between this block and the previous block
			dist, _ := hamming.StringDistance(string(previousKeySizeBlock), string(keySizeBlock))

			totalKeySizeBlockDist += dist
			numKeySizeBlocksUsed++
		}
		averageDist := float64(totalKeySizeBlockDist) / float64(numKeySizeBlocksUsed)
		normalizedDist := averageDist / float64(keySize)

		for i := 0; i < len(minNormalizedDists); i++ {
			if normalizedDist < minNormalizedDists[i] {
				// if this normalizedDist is less then a value in the minNormalizedDists array, then insert it before that item,
				// and shift the rest of the array down
				copy(minNormalizedDists[i+1:], minNormalizedDists[i:])
				minNormalizedDists[i] = normalizedDist
				copy(minNormalizedDistKeySizes[i+1:], minNormalizedDistKeySizes[i:])
				minNormalizedDistKeySizes[i] = keySize

				// don't want to insert this new smallest dist multiple times
				break
			}
		}
	}

	// truncate the probable keysizes so that zeroes aren't included
	probableKeySizes := minNormalizedDistKeySizes
	for i, keySize := range minNormalizedDistKeySizes {
		if keySize == 0 {
			probableKeySizes = minNormalizedDistKeySizes[:i]
		}
	}

	return probableKeySizes
}
