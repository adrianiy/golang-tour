// Package robotname contains methods to manage factory robots
package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot type here.
type Robot struct {
	ID string
}

var usedNames = map[string]bool{}
var possibilities = 26 * 26 * 1000

func getName() string {
	rChar := 'A' + rune(rand.Intn(26))
	rChar2 := 'A' + rune(rand.Intn(26))
	rDigit := rand.Intn(1000)

	return fmt.Sprintf("%s%s%03d", string(rChar), string(rChar2), rDigit)
}

// Name generates robot name taking int account repeated names
func (r *Robot) Name() (string, error) {
	if r.ID != "" {
		return r.ID, nil
	}

	if possibilities == 0 {
		return "", errors.New("All names has been used")
	}

	id := getName()

	if usedNames[id] == true {
		return r.Name()
	}

	usedNames[id] = true
	r.ID = id
	possibilities--

	return id, nil
}

// Reset sets new robot name
func (r *Robot) Reset() {
	r.ID = ""
}
