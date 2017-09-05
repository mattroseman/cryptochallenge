package xorcipher

import (
	"testing"
)

func TestFindEncryptedMessage(t *testing.T) {
	unencryptedLine := FindEncryptedMessage("encryptedMessages.txt")
	result := "Now that the party is jumping"
	if unencryptedLine != result {
		t.Errorf("Finding the unencrypted message found: %s, but expected: %s", unencryptedLine, result)
	}
}
