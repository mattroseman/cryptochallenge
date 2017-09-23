package galois

const primitive = uint16(0x11d)

// reduce takes a binary number that is usually greater than 8 bits, and reduces it by dividing by the
// primitive 0x11d, and getting the remainder.
// @param a: the number to be reduced to 8 bits
// @return: the remainder of a / primitive in GF(2**8)
func reduce(a uint16) byte {
	if getMSBPosition(a) <= 8 {
		return byte(a)
	}
	// the divisor is the primitive shifted over so that the primitive's MSB and a's MSB are lined up
	// xor a and the divisor, then reshift the primitive over so MSBs line up again
	// continue doing this until a <= 8 bits long
	for divisor := primitive << (getMSBPosition(a) - 9); getMSBPosition(a) > 8; divisor = primitive << (getMSBPosition(a) - 9) {
		a ^= divisor
	}

	return byte(a)
}

// getMSBPosition takes a binary number and returns the "length" of the nubmer or the position of it's msb, with
// 1 returning 1, and 100011101 returning 9
func getMSBPosition(a uint16) (bits uint16) {
	for (a >> bits) > 0 {
		bits++
	}

	return
}
