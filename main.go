package main

import (
	"flag"
	"fmt"
	"interview/planets/prediction"
	"interview/planets/solarsystem"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

// var maxPerimeter map[int]float64
var loggerLevels = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
}

var opts struct {
	logger struct {
		level string
	}
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {

}

func main() {

	flag.StringVar(&opts.logger.level, "log-level", "info", "set Log level")
	flag.Parse()

	loglevel, ok := loggerLevels[opts.logger.level]
	if !ok {
		log.Fatalf("Invalid log level: %s", opts.logger.level)
	}
	handleOption := &slog.HandlerOptions{Level: loglevel}
	logger := slog.New(slog.NewTextHandler(os.Stdout, handleOption))

	p1 := solarsystem.NewPlanet("Ferengi", 500, 1, 1)
	p2 := solarsystem.NewPlanet("Betasoide", 2000, 3, 1)
	p3 := solarsystem.NewPlanet("Vulcano", 1000, 5, -1)

	//maxPerimeter := map[int]float64{}
	s := solarsystem.SolarSystem{p1, p2, p3}
	//fmt.Println("%+v", s)
	var clima string
	//http.HandleFunc("/clima")

	mux := http.NewServeMux()
	mux.HandleFunc("/clima", func(w http.ResponseWriter, r *http.Request) {

		logger.Info("inside clima endpoint")
		w.Header().Set("Content-Type", "application/json")
		dia := r.URL.Query().Get("dia")
		day, err := strconv.Atoi(dia)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		s.NewPosition(day)
		if prediction.Sequia(s) {
			clima = "sequia"
		}
		if prediction.CondicionesOptimas(s) {
			clima = "condiciones optimas"
		}
		switch {
		case prediction.Sequia(s):
			clima = "sequia"
		case prediction.CondicionesOptimas(s):
			clima = "condiciones optimas"
		default:
			clima = "lluvia"
		}

		jsonResponse := `{"dia":%d, "clima":%s}`
		w.Write([]byte(fmt.Sprintf(jsonResponse, day, clima)))
		//fmt.Fprintf(w, " day %d clima %s", day, clima)
	})
	logger.Info("starting Api-Rest")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}

	//maxPerDays := s.MaxPerimeterDays(10 * 365)
	//fmt.Println(maxPerDays)
	//days := 365 * 1
	//var count int
	//for d := 0; d < days; d++ {
	//	s.NewPosition(d)
	//	fmt.Printf(" %+v %+v %+v \n", s[0], s[1], s[2])
	//
	//	//if prediction.Sequia(s) {
	//	if prediction.CondicionesOptimas(s) {
	//		count++
	//		fmt.Printf("###################a %+v %+v %+v \n", s[0], s[1], s[2])
	//	}
	//}
	//fmt.Println(count)
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
