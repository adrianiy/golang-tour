// Package romannumerals contains solutions for exercism roman numerals: https://exercism.org/tracks/go/exercises/roman-numerals.
package romannumerals

import (
	"errors"
)


var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var letters = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

// ToRomanNumeral function return a given number in roman format or an error.
func ToRomanNumeral(input int) (string, error) {
    var roman string

    if input > 3000 || input <= 0 {
        return roman, errors.New("invalid number")
    }

    for input > 0 {
        for i, value := range values {
            letter := letters[i]
            if input >= value {
                roman += letter
                input -= value
                break
            }
        }
    }

    return roman, nil
    
}
