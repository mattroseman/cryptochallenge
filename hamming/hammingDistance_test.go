package hamming

import "testing"

func TestStringDistance(t *testing.T) {
	string1 := "this is a test"
	string2 := "wokka wokka!!!"
	dist, err := StringDistance(string1, string2)
	if err != nil {
		t.Errorf("HammingDistance was incorrect, raised error: %s", err)
	}
	expectedResult := 37
	if dist != expectedResult {
		t.Errorf("HammingDistance between \"this is a test\" and \"wokka wokka!!!\" expected %d, but got %d", expectedResult, dist)
	}
}
