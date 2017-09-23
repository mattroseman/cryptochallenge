package galois

// Sub  takes two bytes and subtracts them in GF(2**8) which is the same as addition
func Sub(a, b byte) byte {
	return Add(a, b)
}
