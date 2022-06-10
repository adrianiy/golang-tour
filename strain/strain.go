// Package strain contains solutions for exercism strain: https://exercism.org/tracks/go/exercises/strain.
package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep function returns those elements that match predicate condition.
func (i Ints) Keep(filter func(int) bool) Ints {
    var result Ints

    for _, element := range i {
        if filter(element) {
            result = append(result, element)
        }
    }

    return result
}

// Discard function discard those elements that match predicate condition.
func (i Ints) Discard(filter func(int) bool) Ints {
    var result Ints

    for _, element := range i {
        if !filter(element) {
            result = append(result, element)
        }
    }

    return result
}

// Keep function returns lists that match predicate condition.
func (l Lists) Keep(filter func([]int) bool) Lists {
    var result Lists 

    for _, list := range l {
        if filter(list) {
            result = append(result, list)
        }
    }

    return result
}

// Keep function returns strings that match predicate condition.
func (s Strings) Keep(filter func(string) bool) Strings {
    var result Strings

    for _, element := range s {
        if filter(element) {
            result = append(result, element)
        }
    }

    return result
}
