// Package raindrops has a function which returns a sound string
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns a sound string for factors of 3,5,7
func Convert(number int) string {
	var stringBuilder strings.Builder
	if number%3 == 0 {
		stringBuilder.WriteString("Pling")
	}
	if number%5 == 0 {
		stringBuilder.WriteString("Plang")
	}
	if number%7 == 0 {
		stringBuilder.WriteString("Plong")
	}
	if stringBuilder.String() == "" {
		return strconv.Itoa(number)
	}
	return stringBuilder.String()
}
