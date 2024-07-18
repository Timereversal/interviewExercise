package prediction

import (
	"fmt"
	"interview/planets/solarsystem"
	"math"
)

const epsilon = 1e-2

// Sequia return a boolean. if planets are aligned with the sun (return true), solar system weather is sequia
func Sequia(s solarsystem.Planets) bool {

	if solarsystem.GetAngle(*s[0], *s[1])%180 == 0 && solarsystem.GetAngle(*s[1], *s[2])%180 == 0 {

		return true
	}
	return false
}

// lluviav returns 2 booleans, first boolean reports the lluvia state, second boolean report lluvia intensity.
// lluvia state is true when the sun is inside a triangle composed by the 3 planets.
// lluvia intensity is true when the triangle's perimeter has the highest value.
func lluvia(s solarsystem.Planets, day int) (bool, bool) {
	lluvias := false
	intensity := false

	if Sequia(s) {
		return false, false
	}

	if solarsystem.GetAngle(*s[0], *s[1])+solarsystem.GetAngle(*s[1], *s[2])+solarsystem.GetAngle(*s[2], *s[0]) == 360 {
		lluvias = true
	}
	for _, v := range s.MaxPerimeterDays(365 * 10) {
		if day == v {
			intensity = true
		}
	}

	return lluvias, intensity
}

// slope is the angle's tangent made of the line of Planet 1 and Planet 2 with the X axis.
func slope(p1, p2 solarsystem.Planet) float64 {
	x1 := p1.Radius * math.Cos(float64(p1.Clockwise)*float64(p1.Angle)*math.Pi/180)
	x2 := p2.Radius * math.Cos(float64(p2.Clockwise)*float64(p2.Angle)*math.Pi/180)

	y1 := p1.Radius * math.Sin(float64(p1.Clockwise)*float64(p1.Angle)*math.Pi/180)
	y2 := p2.Radius * math.Sin(float64(p2.Clockwise)*float64(p2.Angle)*math.Pi/180)

	return math.Atan((y2 - y1) / (x2 - x1))
}

func area(s solarsystem.Planets) float64 {
	area := 0.5*math.Abs(s[0].Radius*s[1].Radius)*math.Sin(float64(s[1].Clockwise*s[1].Angle-s[0].Clockwise*s[0].Angle)*math.Pi/180) +
		0.5*math.Abs(s[1].Radius*s[2].Radius)*math.Sin(float64(s[2].Clockwise*s[2].Angle-s[1].Clockwise*s[1].Angle)*math.Pi/180) -
		0.5*math.Abs(s[0].Radius*s[2].Radius)*math.Sin(float64(s[2].Clockwise*s[2].Angle-s[0].Clockwise*s[0].Angle)*math.Pi/180)
	return area
}

// colineal reports whether the planets are on the same line.
// it returns true if the slope of line of planet 1 - planet 2 is equal(delta 0.01) to line's slope of planet 2 - planet 3
// due to float comparison, I am using equal function , which has an error of 0.01.
func colineal(s solarsystem.Planets) bool {
	if equal(slope(*s[0], *s[1]), slope(*s[1], *s[2])) {
		return true
	}
	return false
}

// CondicionesOptimas reports whether the solarsystem is in Optimal Conditions.
// Optimal conditions means, planets are colineal without the sun
func CondicionesOptimas(s solarsystem.Planets) bool {

	if Sequia(s) {
		fmt.Println("sequia")
		return false

	}
	a := colineal(s)
	return a
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

// Clima report the weather of solarsystem as a string value.
// clima posible values are sequia, condiciones optimas, lluvia.
func Clima(solarsystem solarsystem.Planets, day int) (clima string) {
	solarsystem.NewPosition(day)
	switch {
	case Sequia(solarsystem):
		clima = "sequia"
	case CondicionesOptimas(solarsystem):
		clima = "condiciones optimas"
	default:
		clima = "lluvia"

	}
	return clima
}
