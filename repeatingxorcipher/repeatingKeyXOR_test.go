package repeatingxorcipher

import "testing"

func TestRepeatingKeyXOREncrypt(t *testing.T) {
	key := []byte("ICE")
	plaintext := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	encryptedMessage := RepeatingKeyXOREncrypt(plaintext, key)
	expectedResult := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a2622632427276" +
		"5272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if encryptedMessage != expectedResult {
		t.Errorf("RepeatingKeyXOREncrypt returned the wrong result: exptected %s, got %s", expectedResult, encryptedMessage)
	}
}
