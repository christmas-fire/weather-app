package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/christmas-fire/weather-app/config"
	"github.com/christmas-fire/weather-app/internal/models"
	"github.com/christmas-fire/weather-app/internal/out"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

// Create an app
func NewApp() *cli.App {
	app := &cli.App{
		Name:        "weather",
		Usage:       "shows a current weather",
		Description: "This application fetches and displays the current weather for a specified city using the OpenWeatherMap API.\nWithout any args the city sets as 'Novosibirsk'",
		ArgsUsage:   "[city]",
		Action: func(c *cli.Context) error {

			if err := config.InitConfig(); err != nil {
				return fmt.Errorf("error: init config: %s", err.Error())
			}

			apiKey := viper.GetString("api_key")
			city := c.Args().First()
			if city == "" {
				city = viper.GetString("city")
			}
			lang := viper.GetString("lang")

			res, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric&lang=%s", city, apiKey, lang))
			if err != nil {
				return fmt.Errorf("error: failed to make request to weather API: %s", err.Error())
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				if res.StatusCode == 404 {
					return fmt.Errorf("error: city not found")
				}
				return fmt.Errorf("error: returned status code %d", res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				return fmt.Errorf("error: reading request body: %s", err.Error())
			}

			var weatherData models.WeatherData

			err = json.Unmarshal(body, &weatherData)
			if err != nil {
				return fmt.Errorf("error: unmarshal json: %s", err.Error())
			}

			out.PrintWeatherData(weatherData, lang)

			return nil
		},
	}
	return app
}
