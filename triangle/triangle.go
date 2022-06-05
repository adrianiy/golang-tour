// Package triangle contains solutions for traingle exercise
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	NaT = 0 // not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// isNat function returns true is sides does not form a triangle.
func isNaT(a, b, c float64) bool {
	invalidSides := a <= 0 || b <= 0 || c <= 0
	invalidShape := a+b < c || a+c < b || b+c < a

	return invalidSides || invalidShape
}

// isEqu returns true if traingle is equilateral.
func isEqu(a, b, c float64) bool {
	return a == b && b == c
}

// isSca function returns true if triangle is scalene.
func isSca(a, b, c float64) bool {
	return a != b && b != c && a != c
}

// KindFromSides function return the type of triangle based on the side sizes
func KindFromSides(a, b, c float64) Kind {
	if isNaT(a, b, c) {
		return NaT
	} else if isEqu(a, b, c) {
		return Equ
	} else if isSca(a, b, c) {
		return Sca
	} else {
		return Iso
	}
}
