package galois

import "math"

// Pow calculates the power of a to pow in GF(2**8)
func Pow(a byte, pow int) byte {
	return expTable[int(math.Mod(float64(logTable[a])*float64(pow), 255))]
}
