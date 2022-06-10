// Package dna contains response for exercise dna https://exercism.org/tracks/go/exercises/nucleotide-count
package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (h Histogram) count(letter rune) (e error) {
    if _, ok := h[letter]; ok {
        h[letter]++
    } else {
        e = errors.New("invalid chain")
    }
    return
}

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (Histogram, error) {
    var h  = Histogram{ 'A': 0, 'C': 0, 'G': 0, 'T': 0 }

    for _, letter := range string(d) {
        err := h.count(letter)

        if err != nil {
            return h, err
        }
    }
	return h, nil
}
