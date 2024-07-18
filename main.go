package main

import (
	"flag"
	"fmt"
	"interview/planets/controllers"
	"interview/planets/solarsystem"
	"log"
	"log/slog"
	"net/http"
	"os"
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

func NewServer() http.Handler {
	mux := http.NewServeMux()
	return mux
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

	s := solarsystem.Planets{p1, p2, p3}

	ss := controllers.SolarSystem{Solar: s}
	mux := http.NewServeMux()
	mux.HandleFunc("/clima", ss.Clima)
	logger.Info("starting Api-Rest")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}

}
