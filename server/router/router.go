package router

import (
	"database/sql"
	"net/http"
	"y/handler"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router(db *sql.DB) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/weather", handler.GetWeatherDataHandler(db)).Methods("GET", "OPTIONS")

	router.HandleFunc("/api/v1/weather/chart", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWeatherChart(db, w, r)
	}).Methods("GET", "OPTIONS")

	
	return router
}
