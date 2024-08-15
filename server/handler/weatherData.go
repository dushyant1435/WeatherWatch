package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"y/models"

	_ "github.com/lib/pq"
)

var cities = []string{"Delhi", "Mumbai", "Chennai", "Bangalore", "Kolkata", "Hyderabad"}

// FetchAndStoreWeatherData fetches weather data for multiple cities and stores it in the database every 5 minutes
func FetchAndStoreWeatherData(db *sql.DB) {
	for {
		for _, city := range cities {
			data, err := fetchWeatherData(city)
			if err != nil {
				log.Printf("Error fetching data for city %s: %v", city, err)
				continue
			}

			err = storeWeatherData(db, data)
			if err != nil {
				log.Printf("Error storing data for city %s: %v", city, err)
				continue
			}
		}

		time.Sleep(8 * time.Second) // Wait for 5 minutes before fetching data again
	}
}

// fetchWeatherData retrieves weather data from the OpenWeatherMap API for a specific city
func fetchWeatherData(city string) (models.WeatherData, error) {
	var weatherData models.WeatherData

	apiKey := "dc618f9673d2016ade9c2f1ac7ac668c" // Replace with your actual API key
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return weatherData, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return weatherData, err
	}

	main := result["main"].(map[string]interface{})
	weather := result["weather"].([]interface{})[0].(map[string]interface{})

	weatherData.CityName = city
	weatherData.Timestamp = time.Unix(int64(result["dt"].(float64)), 0).Format(time.RFC3339)
	weatherData.Temperature = main["temp"].(float64) - 273.15 // Convert from Kelvin to Celsius
	weatherData.FeelsLike = main["feels_like"].(float64) - 273.15
	weatherData.WeatherMain = weather["main"].(string)

	return weatherData, nil
}

// storeWeatherData inserts the weather data into the database
func storeWeatherData(db *sql.DB, data models.WeatherData) error {
	sqlStatement := `
        INSERT INTO weather_data (city_name, timestamp, temperature, feels_like, weather_main)
        VALUES ($1, $2, $3, $4, $5)`

	_, err := db.Exec(sqlStatement, data.CityName, data.Timestamp, data.Temperature, data.FeelsLike, data.WeatherMain)
	if err != nil {
		return err
	}

	// After storing data, check recent data for alert conditions
	err = CheckRecentWeatherData(db, data.CityName)
	if err != nil {
		log.Printf("Error checking recent weather data: %v", err)
	}

	return nil
}

// GetWeatherDataHandler retrieves weather data for a specific city and date from the database
func GetWeatherDataHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		city := queryParams.Get("city")
		date := queryParams.Get("date") // Expected format: "YYYY-MM-DD"

		if city == "" || date == "" {
			http.Error(w, "City and date parameters are required", http.StatusBadRequest)
			return
		}

		sqlStatement := `
			SELECT id, city_name, timestamp, temperature, feels_like, weather_main
			FROM weather_data
			WHERE city_name = $1 AND DATE(timestamp) = $2`

		rows, err := db.Query(sqlStatement, city, date)
		if err != nil {
			http.Error(w, fmt.Sprintf("Unable to execute the query: %v", err), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var weatherData []models.WeatherData

		for rows.Next() {
			var data models.WeatherData
			err := rows.Scan(&data.ID, &data.CityName, &data.Timestamp, &data.Temperature, &data.FeelsLike, &data.WeatherMain)
			if err != nil {
				http.Error(w, fmt.Sprintf("Unable to scan the row: %v", err), http.StatusInternalServerError)
				return
			}
			weatherData = append(weatherData, data)
		}

		if err := rows.Err(); err != nil {
			http.Error(w, fmt.Sprintf("Row iteration error: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(weatherData)
	}
}
