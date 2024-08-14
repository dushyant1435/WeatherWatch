package handler

import (
	"database/sql"
	"time"
)

// Function to schedule daily weather summary calculation
func ScheduleDailyWeatherSummary(db *sql.DB) {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		cities := []string{"Delhi", "Mumbai", "Chennai", "Bangalore", "Kolkata", "Hyderabad"}
		for _, city := range cities {
			CalculateDailyWeatherSummary(db, city, time.Now().AddDate(0, 0, -1)) // Calculate for yesterday's data
		}
	}
}
