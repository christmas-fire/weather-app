package out

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/christmas-fire/weather-app/internal/models"
)

// Print a formatted weather data
func PrintWeatherData(wd models.WeatherData, lang string) {
	dt := time.Unix(wd.Date, 0)
	date := dt.Format("02-01-2006")

	city := wd.Name

	temp := wd.Main.Temp
	feels_like := wd.Main.Feels_like

	description := wd.Weather[0].Description

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#248AFD")).
		Padding(1, 2)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#248AFD")).
		MarginBottom(1).Align(lipgloss.Center)

	const (
		dateWidth    = 12
		cityWidth    = 15
		weatherWidth = 30
		tempWidth    = 13
		feelsWidth   = 16
	)

	itemStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Center)

	headerStyle := itemStyle.Copy().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#248AFD"))

	dateStyle := itemStyle.Copy().
		Foreground(lipgloss.Color("#888888")).
		Width(dateWidth)

	cityStyle := itemStyle.Copy().
		Foreground(lipgloss.Color("#00FFFF")).
		Width(cityWidth)

	weatherStyle := itemStyle.Copy().
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(weatherWidth)

	tempStyle := itemStyle.Copy().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true).
		Width(tempWidth)

	feelsStyle := itemStyle.Copy().
		Foreground(lipgloss.Color("#FF69B4")).
		Width(feelsWidth)

	var headerDate, headerCity, headerWeather, headerTemp, headerFeels string
	title := "WEATHER APP"
	if lang == "ru" || lang == "russian" {
		headerDate = "Дата"
		headerCity = "Город"
		headerWeather = "Погода"
		headerTemp = "Температура"
		headerFeels = "Ощущается как"
	} else {
		headerDate = "Date"
		headerCity = "City"
		headerWeather = "Weather"
		headerTemp = "Temperature"
		headerFeels = "Feels like"
	}

	headerRow := lipgloss.JoinHorizontal(lipgloss.Top,
		headerStyle.Copy().Width(dateWidth).Render(headerDate),
		headerStyle.Copy().Width(cityWidth).Render(headerCity),
		headerStyle.Copy().Width(weatherWidth).Render(headerWeather),
		headerStyle.Copy().Width(tempWidth).Render(headerTemp),
		headerStyle.Copy().Width(feelsWidth).Render(headerFeels),
	)

	dataRow := lipgloss.JoinHorizontal(lipgloss.Top,
		dateStyle.Render(date),
		cityStyle.Render(city),
		weatherStyle.Render(description),
		tempStyle.Render(fmt.Sprintf("%.0f°", temp)),
		feelsStyle.Render(fmt.Sprintf("%.0f°", feels_like)),
	)

	content := lipgloss.JoinVertical(lipgloss.Center,
		titleStyle.Render(title),
		headerRow,
		dataRow,
	)

	fmt.Println(borderStyle.Align(lipgloss.Center).Render(content))
}
