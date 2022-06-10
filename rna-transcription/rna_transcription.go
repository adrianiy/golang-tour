// Package strand contains solutions for rna exercise: https://exercism.org/tracks/go/exercises/rna-transcription.
package strand

import "fmt"

func transcribe(s rune) string {
    switch(s) {
    case 'G':
        return "C"
    case 'C':
        return "G"
    case 'T':
        return "A";
    default:
        return "U"
    }

}
// ToRNA function returns rna transcription string.
func ToRNA(dna string) string {
    var rna string

    for _, letter := range dna {
        rna = fmt.Sprintf("%s%s", rna, transcribe(letter))
    }

    return rna
}
