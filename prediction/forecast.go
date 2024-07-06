package prediction

import (
	"math"
)

type Planet struct {
	name             string
	radius           float64
	angleSpeedPerDay int
	angle            int
	clockwise        bool
}

func NewSystem(p1, p2, p3 Planet, d int) {

}

// alineados con el sol - then sequia
func sequia(p1, p2, p3 Planet) bool {
	//
	if (math.Tan(float64(p1.angle)) == math.Tan(float64(p2.angle))) && (math.Tan(float64(p2.angle)) == math.Tan(float64(p3.angle))) {
		return true
	}
	return false
}

func lluvia(p1, p2, p3 Planet) (bool, bool) {
	var lluvia bool
	var intensity bool
	if getAngle(p1, p2)+getAngle(p2, p3)+getAngle(p3, p1) == 360 {
		lluvia = true
	}
	lluvia = false
	return lluvia, intensity
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getAngle(p1, p2 Planet) int {

	a := p1.angle % 360
	b := p2.angle % 360

	if p1.clockwise == p2.clockwise {
		angle := absInt(a - b)
		if angle > 180 {
			return 180 - angle
		}
		return angle

	}
	angle := absInt(360 - a + b)
	if angle > 180 {
		return 180 - angle
	}
	return angle

}

func distance(p1, p2 Planet) float64 {
	d := math.Sqrt(math.Pow(p1.radius, 2) + math.Pow(p2.radius, 2) - 2*p1.radius*p2.radius*math.Cos(float64(getAngle(p1, p2))*math.Pi/180))
	return d
}

func perimeter(p1, p2, p3 Planet) float64 {
	var p float64

	p = distance(p1, p2) + distance(p1, p3) + distance(p2, p3)
	return p
}
