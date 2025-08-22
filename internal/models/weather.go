package models

// WeatherData represents the structure of the weather data received from the API.
type WeatherData struct {
	Name string `json:"name"` // example: "London"
	Date int64  `json:"dt"`   // example: 1755413914
	Main struct {
		Temp       float64 `json:"temp"`       // example: 15.42
		Feels_like float64 `json:"feels_like"` // example: 15
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"` // example: "clear sky"
	} `json:"weather"`
}
