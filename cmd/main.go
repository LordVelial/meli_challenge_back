package main

import (
	"fmt"
	"log"
	"meli_challenge_back/cmd/config"
	"meli_challenge_back/internal/handler"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	cfgDb, err := config.GetConfig(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := mux.NewRouter()

	// Rutas para crear y obtener eventos
	router.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateEvent(w, r, cfgDb)
	}).Methods(http.MethodPost)

	router.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEvents(w, r, cfgDb)
	}).Methods(http.MethodGet)

	// Ruta para obtener la lista de países
	router.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		handler.GetCountries(w, r, cfgDb)
	}).Methods(http.MethodGet)

	// Ruta para obtener la lista de tipos de evento
	router.HandleFunc("/event-types", func(w http.ResponseWriter, r *http.Request) {
		handler.GetEventTypes(w, r, cfgDb)
	}).Methods(http.MethodGet)

	// Ruta para obtener las métricas de los países con mayor cantidad de eventos en el último mes
	router.HandleFunc("/metrics/top-countries", func(w http.ResponseWriter, r *http.Request) {
		handler.GetTopCountriesMetrics(w, r, cfgDb)
	}).Methods(http.MethodGet)

	// Crear un manejador CORS
	c := cors.Default()

	// Envolver el enrutador con el manejador CORS
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
