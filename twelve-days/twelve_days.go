package twelve

import "fmt"

var ordinals = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

func getVerseElement(i int) string {
	switch i {
	case 0:
		return "a Partridge in a Pear Tree"
	case 1:
		return "two Turtle Doves"
	case 2:
		return "three French Hens"
	case 3:
		return "four Calling Birds"
	case 4:
		return "five Gold Rings"
	case 5:
		return "six Geese-a-Laying"
	case 6:
		return "seven Swans-a-Swimming"
	case 7:
		return "eight Maids-a-Milking"
	case 8:
		return "nine Ladies Dancing"
	case 9:
		return "ten Lords-a-Leaping"
	case 10:
		return "eleven Pipers Piping"
	default:
		return "twelve Drummers Drumming"
	}
}

func Verse(i int) string {
	i--
	var verse string
	j := 0
	for j <= i {
		join := ", "

		if j == i {
			join = ", and "
		}
		if j == 0 {
			join = ""
		}

		verse += join + getVerseElement(i-j)
		j++
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", ordinals[i], verse)

}

func Song() string {
	var song string
	for i := range [12]int{} {
		breakLine := "\n"

		if i == 11 {
			breakLine = ""
		}

		song += fmt.Sprintf("%s%s", Verse(i+1), breakLine)

	}

	return song
}
