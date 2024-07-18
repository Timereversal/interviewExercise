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

	s := solarsystem.SolarSystem{p1, p2, p3}

	var clima string

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

}
