package hamming

import "errors"

// StringDistance computes the hamming distance between two strings
// @param string1: a string that may be similar to string2
// @param string2: a string that may be similar to string1
// @return: an integer representing the hamming distance between string1 and string2 and an error value
func StringDistance(string1, string2 string) (int, error) {
	// check that the strings are of the same length
	if len(string1) != len(string2) {
		return 0, errors.New("string1 and string2 must be the same length")
	}

	// convert the strings to byte arrays
	bytes1 := []byte(string1)
	bytes2 := []byte(string2)

	var diff int

	for i := 0; i < len(bytes1); i++ {
		byte1 := bytes1[i]
		byte2 := bytes2[i]

		diff += Distance(byte1, byte2)
	}

	return diff, nil
}

// Distance calculates the hamming distance between two bytes
// @param byte1: the first byte
// @param byte2: the second byte
// @return: the hamming distance between these two bytes
func Distance(byte1, byte2 byte) int {
	var diff int

	for j := 0; j < 8; j++ {
		mask := byte(1 << uint(j))
		if (byte1 & mask) != (byte2 & mask) {
			diff++
		}
	}

	return diff
}
