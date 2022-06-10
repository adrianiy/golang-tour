// Package proverb contain result of exercise proverb of exercism https://exercism.org/tracks/go/exercises/proverb
package proverb

import "fmt"

// Proverb function returns the provebial rhyme given a list of strings
func Proverb(rhyme []string) []string {
    var proverb []string

    for i, word := range rhyme {
        var phrase string

        if i < len(rhyme) - 1 {
            phrase = fmt.Sprintf("For want of a %s the %s was lost.", word, rhyme[i+1])
        } else {
            phrase = fmt.Sprintf("And all for the want of a %s.", rhyme[0])
        }

        proverb = append(proverb, phrase)
    }

    return proverb
}
