package xorcipher

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/mroseman95/cryptochallenge/letterdistribution"
)

// FindEncryptedMessage takes a filename, where every line is a hex encoded encrypted message, and
// finds the line that contains a secret message
// @param fileName: the name of the file to read
// @return: the secret message that was found
func FindEncryptedMessage(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// minUnencryptedDiff is the smallest difference gotten by unencrypting the file
	minUnencryptedDiff := math.Inf(1)
	// minUnencryptedSentance is the sentance that has the minimum difference from normal english text
	var minUnencryptedSentance string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hexLine := scanner.Text()

		// cut off the newline
		hexLine = hexLine[:len(hexLine)-2]

		line := UnencryptXOREncryptedHexString(hexLine)
		lineBytes := []byte(line)

		// find the most likely unencrypted value of this line
		_, unencryptedText := FindXORByte(lineBytes)
		// get the diff of the unencrypted line with normal english text
		diff := letterdistribution.GetEnglishCharacterProbability(unencryptedText)
		if diff < minUnencryptedDiff {
			minUnencryptedSentance = unencryptedText
			minUnencryptedDiff = diff
		}
	}

	return minUnencryptedSentance
}
