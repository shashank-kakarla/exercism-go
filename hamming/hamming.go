// Package hamming contains a function which returns the Hamming Distance in the 2 strings
package hamming

import (
	"errors"
)

// Distance function calculates the Hamming Distance in the 2 strings
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("different length")
	}

	runeA := []rune(a)
	runeB := []rune(b)
	var length int

	for i := range runeA {
		if runeA[i] != runeB[i] {
			length++
		}
	}

	return length, nil
}
