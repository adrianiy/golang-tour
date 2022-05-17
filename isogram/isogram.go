// Package isogram expose a function to determine if a word is an isogram or not.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram creates a map with letter repetition and return true if the maximum value is 1
func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i, letter := range word {
		if unicode.IsLetter(letter) && strings.ContainsRune(word[i+1:], letter) {
			return false
		}

	}
	return true
}
