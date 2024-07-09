package solarsystem

import (
	"math"
)

type Planet struct {
	Name             string
	Radius           float64
	AngleSpeedPerDay int
	Angle            int
	Clockwise        int
}

type SolarSystem []*Planet

func NewPlanet(name string, radius float64, angleSpeedPerDay int, clockwise int) *Planet {
	return &Planet{
		Name:             name,
		Radius:           radius,
		AngleSpeedPerDay: angleSpeedPerDay,
		Clockwise:        clockwise,
	}
}

func (s SolarSystem) NewPosition(day int) {
	for _, planet := range s {
		planet.Angle = planet.AngleSpeedPerDay * day
	}
}

func (s SolarSystem) MaxPerimeterDays(days int) []int {
	maxPerimeter := map[int]float64{}
	//daysMaxPerimeter := []int{}
	var daysMaxPerimeter []int
	var permax float64
	for d := 1; d < days; d++ {
		s.NewPosition(d)
		perimeter := Perimeter(*s[0], *s[1], *s[2])
		//fmt.Printf("Perimeter: %.2f day: %d p1 %+v distance d1 %f d2 %f d3 %f - angle p1,p2 %d \n", perimeter, d, *s[0], prediction.Distance(*s[0], *s[1]), prediction.Distance(*s[1], *s[2]), prediction.Distance(*s[2], *s[0]), prediction.GetAngle(*s[1], *s[2]))
		if perimeter >= permax {
			permax = perimeter
			maxPerimeter[d] = permax
			//fmt.Println("Perimeter:", permax)
			//fmt.Printf("Perimeter: %.2f\n day: %d ", permax, d)
		}
	}
	for day, perimeter := range maxPerimeter {
		if perimeter == permax {
			daysMaxPerimeter = append(daysMaxPerimeter, day)
		}
	}
	return daysMaxPerimeter

}

func Distance(p1, p2 Planet) float64 {
	d := math.Sqrt(math.Pow(p1.Radius, 2) + math.Pow(p2.Radius, 2) - 2*p1.Radius*p2.Radius*math.Cos(float64(GetAngle(p1, p2))*math.Pi/180))
	return d
}

func Perimeter(p1, p2, p3 Planet) float64 {

	p := Distance(p1, p2) + Distance(p1, p3) + Distance(p2, p3)
	return p
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetAngle(p1, p2 Planet) int {

	a := p1.Angle % 360
	b := p2.Angle % 360

	if p1.Clockwise == p2.Clockwise {
		Angle := absInt(a - b)
		if Angle > 180 {
			return 180 - Angle
		}
		return Angle

	}
	Angle := a + b
	if Angle > 180 {
		return 360 - Angle
	}
	return Angle

}
