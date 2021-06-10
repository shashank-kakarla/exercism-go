// Package twofer has a function to include any given input into a string
package twofer

import (
	"fmt"
)

// ShareWith inserts the given input into a string.
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
