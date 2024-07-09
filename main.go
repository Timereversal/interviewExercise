package main

import (
	"fmt"
	"interview/planets/prediction"
	"interview/planets/solarsystem"
)

//var maxPerimeter map[int]float64

func main() {
	p1 := solarsystem.NewPlanet("Ferengi", 500, 1, 1)
	p2 := solarsystem.NewPlanet("Betasoide", 2000, 3, 1)
	p3 := solarsystem.NewPlanet("Vulcano", 1000, 5, -1)

	//maxPerimeter := map[int]float64{}
	s := solarsystem.SolarSystem{p1, p2, p3}
	//maxPerDays := s.MaxPerimeterDays(10 * 365)
	//fmt.Println(maxPerDays)
	days := 365 * 1
	var count int
	for d := 0; d < days; d++ {
		s.NewPosition(d)
		//fmt.Printf(" %+v %+v %+v \n", s[0], s[1], s[2])

		//if prediction.Sequia(s) {
		if prediction.CondicionesOptimas(s) {
			count++
			fmt.Printf("###################a %+v %+v %+v \n", s[0], s[1], s[2])
		}
	}
	fmt.Println(count)
	//var permax float64
	//for d := 1; d < 10*365; d++ {
	//	s.NewPosition(d)
	//	perimeter := prediction.Perimeter(*s[0], *s[1], *s[2])
	//	//fmt.Printf("Perimeter: %.2f day: %d p1 %+v distance d1 %f d2 %f d3 %f - angle p1,p2 %d \n", perimeter, d, *s[0], prediction.Distance(*s[0], *s[1]), prediction.Distance(*s[1], *s[2]), prediction.Distance(*s[2], *s[0]), prediction.GetAngle(*s[1], *s[2]))
	//	if perimeter >= permax {
	//		permax = perimeter
	//		maxPerimeter[d] = permax
	//		//fmt.Println("Perimeter:", permax)
	//		//fmt.Printf("Perimeter: %.2f\n day: %d ", permax, d)
	//	}
	//}
	// 360/5 , 360/1, 360/3
	// 72,360,120
	//fmt.Println(maxPerimeter)
}
