// Package isogram expose a function to determine if a word is an isogram or not.
package isogram

import "strings"

// IsIsogram creates a map with letter repetition and return true if the maximum value is 1
func IsIsogram(word string) bool {
	accum := map[rune]int8{}

	for _, letter := range strings.ToLower(word) {
		if letter != '-' && letter != ' ' {
			accum[letter]++

			if accum[letter] > 1 {
				return false
			}
		}

	}

	return true
}
