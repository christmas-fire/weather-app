package format

import (
	"fmt"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/christmas-fire/weather-app/internal/models"
	"github.com/olekukonko/tablewriter"
)

// Print a table with weather data
func PrintWeatherData(wd models.WeatherData, lang string) {
	dt := time.Unix(wd.Date, 0)
	date := dt.Format("02-01-2006")

	city := wd.Name

	weather_data := strings.Split(fmt.Sprintf("%s", wd.Weather[0]), " ")
	weather := removeChars(weather_data[0], '{', '}')
	emoji := getWeatherEmoji(weather)
	description := concatDescription(weather_data)

	temp := wd.Main.Temp
	feels_like := wd.Main.Feels_like

	table := tablewriter.NewWriter(os.Stdout)
	if lang == "ru" {
		table.SetHeader([]string{"Дата", "Город", "Погода", "Температура", "Ощущается как"})
	}
	table.SetHeader([]string{"Date", "City", "Weather", "Temperature", "Feels like"})
	table.SetCenterSeparator("|")
	table.SetAutoWrapText(false)

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgBlueColor},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.FgHiCyanColor},
		tablewriter.Colors{tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.FgHiWhiteColor},
		tablewriter.Colors{tablewriter.FgHiWhiteColor},
	)

	table.Append([]string{date, city, fmt.Sprintf("%s  %s", emoji, description), fmt.Sprintf("%.0f°", temp), fmt.Sprintf("%.0f°", feels_like)})

	table.Render()
}

// Remove specifed chars from the string
func removeChars(str string, chars ...rune) string {
	for _, char := range chars {
		str = strings.ReplaceAll(str, string(char), "")
	}
	return str
}

// Concatenate parts of JSON weather.description field into string
func concatDescription(weather []string) string {
	res := ""
	for i, part := range weather {
		// weather.main
		if i == 0 {
			continue
		}
		// weather.description
		if i == 1 {
			str := []rune(part)
			str[0] = unicode.ToUpper(str[0])
			part = string(str)
		}
		res += part + " "
	}
	res = removeChars(res, '{', '}')
	return res
}

// Get an emoji depends on JSON weather.main field
func getWeatherEmoji(main string) string {
	switch main {
	case "Thunderstorm":
		return "⛈️"
	case "Drizzle":
		return "🌧️"
	case "Rain":
		return "☔"
	case "Snow":
		return "❄️"
	case "Atmosphere":
		return "🌫️"
	case "Clear":
		return "☀️"
	case "Clouds":
		return "☁️"
	default:
		return ""
	}
}
