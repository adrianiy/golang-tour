//Package clock define some clock methods.
package clock

import "fmt"

//Clock type structure.
type Clock struct {
	Hour   int
	Minute int
}

//New defines a new hour taking into account all possible minutes and hour values.
func New(h, m int) Clock {
	hoursFromMinutes := m / 60
	h += hoursFromMinutes
	h = h % 24
	m = m % 60

	if m < 0 {
		m = 60 + m

		if m > 0 {
			h--
		}
	}
	if h < 0 {
		h = 24 + h
	}

	return Clock{
		Hour:   h,
		Minute: m,
	}
}

//Add increases clock time by m minutes.
func (c Clock) Add(m int) Clock {
	minutes := c.Minute + m

	return New(c.Hour, minutes)
}

//Subtract decreases clock time by m minutes.
func (c Clock) Subtract(m int) Clock {
	minutes := c.Minute - m

	return New(c.Hour, minutes)
}

//String transform Clock into a string with format HH:mm.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}
