package galois

// Add takes two bytes and adds them together in GF(2**8)
func Add(a, b byte) byte {
	return a ^ b
}
