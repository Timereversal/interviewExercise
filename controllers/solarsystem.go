package controllers

import (
	"fmt"
	"interview/planets/prediction"
	"interview/planets/solarsystem"
	"net/http"
	"strconv"
)

type SolarSystem struct {
	Solar solarsystem.SolarSystem
}

func (s *SolarSystem) Clima(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dia := r.URL.Query().Get("dia")
	day, err := strconv.Atoi(dia)
	if err != nil || day < 0 {
		http.Error(w, `{"message":"day parameter must be >= 0"}`, http.StatusBadRequest)
		return
	}
	clima := prediction.Clima(s.Solar, day)

	jsonResponse := `{"dia":%d, "clima":"%s"}`
	w.Write([]byte(fmt.Sprintf(jsonResponse, day, clima)))
}
