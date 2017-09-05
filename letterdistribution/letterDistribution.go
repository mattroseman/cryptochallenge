package letterdistribution

import (
	"math"
	"strings"
)

// these letter frequencies are from www.data-compression.com/english.html
var letterFrequencies = map[int]float64{
	'a': 0.0652, 'b': 0.0124, 'c': 0.0217, 'd': 0.0350, 'e': 0.1041, 'f': 0.0198, 'g': 0.0159, 'h': 0.0493, 'i': 0.0558,
	'j': 0.0009, 'k': 0.0051, 'l': 0.0331, 'm': 0.0202, 'n': 0.0565, 'o': 0.0596, 'p': 0.0138, 'q': 0.0009, 'r': 0.0498,
	's': 0.0516, 't': 0.0729, 'u': 0.0225, 'v': 0.0083, 'w': 0.0171, 'x': 0.0014, 'y': 0.0016, 'z': 0.0008, ' ': 0.1918,
}

const otherFrequency = 0.0129

// GetEnglishCharacterProbability takes a sentance, and finds the frequency of each character, then compares to normal
// english character frequency. It returns the difference
func GetEnglishCharacterProbability(sentance string) float64 {
	totalLetterCount := float64(len(sentance))
	letterCount := make(map[int]float64)
	otherCount := 0.0
	difference := 0.0

	sentance = strings.ToLower(sentance)

	for _, character := range sentance {
		if character >= 'a' && character <= 'z' || character == ' ' {
			letterCount[int(character)]++
		} else {
			otherCount++
		}
	}

	for letter, freq := range letterFrequencies {
		difference += math.Abs((letterCount[letter] / totalLetterCount) - freq)
	}
	difference += math.Abs((otherCount / totalLetterCount) - otherFrequency)

	return difference
}
