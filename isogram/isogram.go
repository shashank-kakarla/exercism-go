// Package isogram has function which identifies if a word is an Isogram.
package isogram

import (
	"strings"
)

// IsIsogram returns true if the word is an Isogram or false if its is not.
func IsIsogram(word string) bool {
	var repeated = make(map[rune]bool)
	for _, r := range strings.ToLower(word) {
		if _, found := repeated[r]; found && r != '-' && r != ' ' {
			return false
		}
		repeated[r] = false
	}
	return true
}
