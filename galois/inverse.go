package galois

// Inverse calculates the GF(2**8) inverse of a
func Inverse(a byte) byte {
	return expTable[255-logTable[a]]
}
