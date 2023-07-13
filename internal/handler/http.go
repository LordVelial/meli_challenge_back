package handler

import (
	"encoding/json"
	"net/http"

	"meli_challenge_back/internal/entity"
	"meli_challenge_back/internal/service"
)

// Handler para crear un nuevo evento
func CreateEvent(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	switch r.Method {
	case http.MethodPost:
		service.CreateEvent(w, r, cfgDb)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler para obtener eventos con paginación y filtros opcionales
func GetEvents(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	switch r.Method {
	case http.MethodGet:
		service.GetEvents(w, r, cfgDb)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handler para obtener la lista de países
func GetCountries(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	countries, err := service.GetCountries(cfgDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(countries)
}

// Handler para obtener la lista de tipos de evento
func GetEventTypes(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	eventTypes, err := service.GetEventTypes(cfgDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(eventTypes)
}

// Handler para obtener las métricas de los países con mayor cantidad de eventos en el último mes
func GetTopCountriesMetrics(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	topCountries, err := service.GetTopCountriesMetrics(cfgDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(topCountries)
}
