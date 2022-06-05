// Package space contains solutions for Space Age problem
package space

// Planet struct defines Planet type
type Planet string

type orbitalPeriod struct {
	name   string
	period float64
}

var earthYearInMs float64 = 31557600

var orbitalPeriods = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age function return the astronaut age in a determinate planet
func Age(seconds float64, planet Planet) float64 {
	earthYears := seconds / earthYearInMs
	return earthYears / orbitalPeriods[planet]
}
