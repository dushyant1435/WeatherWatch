package models

type Problem struct {
	ID           int     `json:"id"`
	UserId       int     `json:"user_id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Constraints  *string `json:"constraints,omitempty"`   // use *string to allow null values
	InputFormat  *string `json:"input_format,omitempty"`  // use *string to allow null values
	OutputFormat *string `json:"output_format,omitempty"` // use *string to allow null values
	// Status       bool  `json:"status`
}

type TestCase struct {
	ID     int    `json:"id"`
	Input  string `json:"input"`
	Output string `json:"output"`
	Sample bool   `json:"sample"`
}

type CodeData struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	UserID int    `json:"user_id"`
}

type CustomCodeData struct {
	Code  string `json:"code"`
	Input string `json:"input"`
}

type RequestBody struct {
	UserID int64 `json:"user_id"`
}

// type WeatherData struct {
// 	Weather []struct {
// 		Main string `json:"main"`
// 	} `json:"weather"`
// 	Main struct {
// 		Temp      float64 `json:"temp"`
// 		FeelsLike float64 `json:"feels_like"`
// 	} `json:"main"`
// 	Dt   int64  `json:"dt"`
// 	Name string `json:"name"`
// }

type WeatherData struct {
	ID          int     `json:"id"`
	CityName    string  `json:"city_name"`
	Timestamp   string  `json:"timestamp"`
	Temperature float64 `json:"temperature"`
	FeelsLike   float64 `json:"feels_like"`
	WeatherMain string  `json:"weather_main"`
}

type DailyWeatherSummary struct {
	City               string  `json:"city"`
	Date               string  `json:"date"`
	AverageTemperature float64 `json:"average_temperature"`
	MaxTemperature     float64 `json:"max_temperature"`
	MinTemperature     float64 `json:"min_temperature"`
	DominantCondition  string  `json:"dominant_condition"`
}
