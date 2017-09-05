package repeatingxorcipher

import (
	"io/ioutil"
	"testing"
)

func TestProbableKeySizes(t *testing.T) {
	file, err := ioutil.ReadFile("repeatingCipherEncryptedMessage.txt")
	if err != nil {
		panic(err)
	}

	probableKeySizes := ProbableKeySizes([]byte(file))

	expectedResult := []int{3, 2, 11}

	for i := range expectedResult {
		if probableKeySizes[i] != expectedResult[i] {
			t.Errorf("ProbableKeySizes returned wrong result. Expected: %v, but Got: %v", expectedResult, probableKeySizes)
		}
	}
}

func TestBreakRepeatingKeyXOR(t *testing.T) {
	// read base64 encoded repeating XOR encrypted file
	file, err := ioutil.ReadFile("repeatingCipherEncryptedMessage.txt")
	if err != nil {
		panic(err)
	}

	decryptedFile, err := BreakRepeatingKeyXOR(string(file))
	if err != nil {
		panic(err)
	}

	expectedResult, err := ioutil.ReadFile("repeatingCipherDecryptedMessage.txt")
	if err != nil {
		panic(err)
	}

	if decryptedFile != string(expectedResult) {
		t.Errorf("BreakRepeatingKeyXOR expected %s, but got %s", string(expectedResult), decryptedFile)
	}
}
