//Package luhn implements luhn algorithm used to validate identification numbers
// such as credic cards, IMEI, etc.
package luhn

import (
	"regexp"
	"strconv"
	"unicode/utf8"
)

func removeWhiteSpaces(s string) string {
	exp := regexp.MustCompile(`\s`)

	return exp.ReplaceAllString(s, "")
}

//Valid return true if id is validated by luhn algorithm.
func Valid(id string) bool {
	id = removeWhiteSpaces(id)

	if len(id) == 1 {
		return false
	}

	nDigits := utf8.RuneCountInString(id)
	sum, _ := strconv.Atoi(string(id[nDigits-1]))
	parity := (nDigits - 2) % 2

	for i := 0; i <= nDigits-2; i++ {
		digit, err := strconv.Atoi(string(id[i]))

		if err != nil {
			return false
		}

		if i%2 == parity {
			digit *= 2
		}
		if digit > 9 {
			digit -= 9
		}

		sum += digit
	}

	return sum%10 == 0
}
