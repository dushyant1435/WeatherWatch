package main

import (
	"fmt"
	"log"
	"net/http"
	"y/handler"
	"y/router"

	_ "github.com/lib/pq"
)

func main() {
	db := handler.CreateConnection()
	defer db.Close()

	go handler.FetchAndStoreWeatherData(db)
	go handler.ScheduleDailyWeatherSummary(db) // Schedule daily summary

	r := router.Router(db)
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
