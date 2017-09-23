package galois

import (
	"errors"
	"math"
)

// Div takes two bytes and divides them in GF(2**8) field
// @param a: the dividend
// @param b: the divisor
func Div(a, b byte) (byte, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by 0")
	}

	if a == 0 {
		return 0, nil
	}

	return expTable[int(math.Mod(float64(logTable[a]-logTable[b])+255, 255))], nil
}
