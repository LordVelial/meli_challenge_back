package service

import (
	"encoding/json"
	"net/http"

	"meli_challenge_back/internal/db"
	"meli_challenge_back/internal/entity"
)

// Servicio para crear un nuevo evento
func CreateEvent(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	var event entity.EventRequest
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = db.CreateEvent(event, cfgDb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Servicio para obtener eventos con paginación y filtros opcionales
func GetEvents(w http.ResponseWriter, r *http.Request, cfgDb entity.DataBase) {
	// Obtiene los parámetros de consulta
	queryValues := r.URL.Query()
	page := queryValues.Get("page")
	pageSize := queryValues.Get("pageSize")
	description := queryValues.Get("description")
	country := queryValues.Get("country")
	typeEvent := queryValues.Get("type")

	var events []entity.EventResponse
	var err error
	whereQuery := ""
	hasWhere := true

	if description != "" {
		// Filtra los eventos por descripción
		if hasWhere {
			whereQuery += " where "
			hasWhere = false
		}
		whereQuery += db.ListEventsByDescription(description, page, pageSize, cfgDb)
	}
	if country != "" {
		// Filtra los eventos por país
		if hasWhere {
			whereQuery += " where "
			hasWhere = false
		} else {
			whereQuery += " and "
		}

		whereQuery += db.ListEventsByCountry(country, page, pageSize, cfgDb)
	}
	if typeEvent != "" {
		// Filtra los eventos por tipo
		if hasWhere {
			whereQuery += " where "
			hasWhere = false
		} else {
			whereQuery += " and "
		}
		whereQuery += db.ListEventsByType(typeEvent, page, pageSize, cfgDb)
	}

	events, err = db.ListEvents(whereQuery, page, pageSize, cfgDb)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devuelve los eventos en formato JSON
	json.NewEncoder(w).Encode(events)
}

// Servicio para obtener los paises
func GetCountries(cfgDb entity.DataBase) ([]entity.Data, error) {
	return db.GetCountries(cfgDb)
}

// Servicio para obtener los tipos de evento
func GetEventTypes(cfgDb entity.DataBase) ([]entity.Data, error) {
	return db.GetEventTypes(cfgDb)
}

// Servicio para obtener las métricas de los países con mayor cantidad de eventos en el último mes
func GetTopCountriesMetrics(cfgDb entity.DataBase) ([]entity.CountryMetric, error) {
	return db.GetTopCountriesMetrics(cfgDb)
}
