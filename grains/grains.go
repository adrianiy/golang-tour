//Package grains exposes functions to know about how many grains do we have
//on the board.
package grains

import (
	"errors"
	"math"
)

//Square tells us how many grains do we have in a specific square.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Square not valid")
	}

	grains := math.Pow(2, float64(number-1))

	return uint64(grains), nil
}

//Total tells the total amount of grains in the table, we can reuse Square funciton
func Total() uint64 {
	var grains uint64

	for i := 0; i < 64; i++ {
		grains += uint64(math.Pow(2, float64(i)))
	}

	return grains
}
