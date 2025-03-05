package models

// Weather data model
type WeatherData struct {
	Name string `json:"name"`
	Date int64  `json:"dt"`
	Main struct {
		Temp       float64 `json:"temp"`
		Feels_like float64 `json:"feels_like"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
}
