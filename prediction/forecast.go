package prediction

import (
	"fmt"
	"interview/planets/solarsystem"
	"math"
)

const epsilon = 1e-2

// alineados con el sol - then sequia
func Sequia(s solarsystem.SolarSystem) bool {
	//
	//if (math.Tan(float64(s[0].Angle)*math.Pi/180) == math.Tan(float64(s[1].Angle))*math.Pi/180) && (math.Tan(float64(s[0].Angle)*math.Pi/180) == math.Tan(float64(s[2].Angle)*math.Pi/180)) {
	//fmt.Println(math.Tan(float64(s[0].Angle)*math.Pi/180), math.Tan(float64(s[1].Angle)*math.Pi/180))
	//if math.Tan(float64(s[0].Angle)*math.Pi/180) == math.Tan(float64(s[1].Angle)*math.Pi/180) {
	if solarsystem.GetAngle(*s[0], *s[1])%180 == 0 && solarsystem.GetAngle(*s[1], *s[2])%180 == 0 {

		return true
	}
	return false
}

// add perimeter
func lluvia(s solarsystem.SolarSystem, day int) (bool, bool) {
	lluvias := false
	intensity := false

	if Sequia(s) {
		return false, false
	}
	// sun inside triangle
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

func slope(p1, p2 solarsystem.Planet) float64 {
	x1 := p1.Radius * math.Cos(float64(p1.Clockwise)*float64(p1.Angle)*math.Pi/180)
	x2 := p2.Radius * math.Cos(float64(p2.Clockwise)*float64(p2.Angle)*math.Pi/180)

	y1 := p1.Radius * math.Sin(float64(p1.Clockwise)*float64(p1.Angle)*math.Pi/180)
	y2 := p2.Radius * math.Sin(float64(p2.Clockwise)*float64(p2.Angle)*math.Pi/180)
	//fmt.Printf("angle %d %2f %2f \n", p1.Angle, float64(p1.Angle)*math.Pi/180, x1)
	return math.Atan((y2 - y1) / (x2 - x1))
}

func area(s solarsystem.SolarSystem) float64 {
	area := 0.5*math.Abs(s[0].Radius*s[1].Radius)*math.Sin(float64(s[1].Clockwise*s[1].Angle-s[0].Clockwise*s[0].Angle)*math.Pi/180) +
		0.5*math.Abs(s[1].Radius*s[2].Radius)*math.Sin(float64(s[2].Clockwise*s[2].Angle-s[1].Clockwise*s[1].Angle)*math.Pi/180) -
		0.5*math.Abs(s[0].Radius*s[2].Radius)*math.Sin(float64(s[2].Clockwise*s[2].Angle-s[0].Clockwise*s[0].Angle)*math.Pi/180)
	return area
}

func colineal(s solarsystem.SolarSystem) bool {
	//fmt.Println("colineal")
	//fmt.Printf("slope angle1 %2f angle2 %2f \n", slope(*s[0], *s[1]), slope(*s[2], *s[0]))
	if equal(slope(*s[0], *s[1]), slope(*s[1], *s[2])) {
		return true
	}
	return false
}

func CondicionesOptimas(s solarsystem.SolarSystem) bool {
	//fmt.Println("prediction condiciones optimas")
	if Sequia(s) {
		fmt.Println("sequia")
		return false

	}
	a := colineal(s)
	//a := equal(area(s), 0)
	return a
}

func equal(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}
