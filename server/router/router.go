package router

import (
	"database/sql"
	"y/handler"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router(db *sql.DB) *mux.Router {

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/weather", handler.GetWeatherDataHandler(db)).Methods("GET", "OPTIONS")

	return router
}
