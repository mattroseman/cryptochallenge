package galois

import "math"

// slowMult takes two bytes and multiplies them in GF(2**8), by manually doing the math
// this is slower than Mult, which uses the exponent table and adds the logarithms
func slowMult(a, b byte) byte {
	x := uint16(a)
	y := uint16(b)
	product := uint16(0)

	for i := uint16(0); (y >> i) > 0; i++ {
		if y&(1<<i) != 0 {
			product ^= x << i
		}
	}

	return reduce(product)
}

// Mult multiplies a and b bytes in GF(2**8)
func Mult(a, b byte) byte {
	if a == 0 || b == 0 {
		return byte(0)
	}

	return expTable[int(math.Mod(float64(logTable[a]+logTable[b]), 255))]
}
