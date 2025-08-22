package out

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/christmas-fire/weather-app/internal/models"
	"github.com/christmas-fire/weather-app/internal/utils"
)

// PrintWeatherData prints the weather data in a styled box.
// It uses lipgloss for styling and colors temperature based on its value.
func PrintWeatherData(wd models.WeatherData) {
	dt := time.Unix(wd.Date, 0)
	date := dt.Format("02-01-2006")

	city := wd.Name

	temp := wd.Main.Temp
	feels_like := wd.Main.Feels_like

	description := utils.CapitalizeFirst(wd.Weather[0].Description)

	borderStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#248AFD")).
		Padding(1, 2)

	const (
		dateWidth    = 12
		cityWidth    = 15
		weatherWidth = 20
		tempWidth    = 13
		feelsWidth   = 16
	)

	itemStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Align(lipgloss.Center)

	headerStyle := itemStyle.
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#248AFD"))

	dateStyle := itemStyle.
		Foreground(lipgloss.Color("#888888")).
		Width(dateWidth)

	cityStyle := itemStyle.
		Foreground(lipgloss.Color("#00FFFF")).
		Width(cityWidth)

	weatherStyle := itemStyle.
		Foreground(lipgloss.Color("#FFFFFF")).
		Width(weatherWidth)

	tempStyle := itemStyle.
		Foreground(lipgloss.Color(utils.GetColorForTemperature(temp))).
		Width(tempWidth)

	feelsStyle := itemStyle.
		Foreground(lipgloss.Color(utils.GetColorForTemperature(temp))).
		Width(feelsWidth)

	headerRow := lipgloss.JoinHorizontal(lipgloss.Top,
		headerStyle.Width(dateWidth).Render("Date"),
		headerStyle.Width(cityWidth).Render("City"),
		headerStyle.Width(weatherWidth).Render("Weather"),
		headerStyle.Width(tempWidth).Render("Temperature"),
		headerStyle.Width(feelsWidth).Render("Feels like"),
	)

	dataRow := lipgloss.JoinHorizontal(lipgloss.Top,
		dateStyle.Render(date),
		cityStyle.Render(city),
		weatherStyle.Render(description),
		tempStyle.Render(fmt.Sprintf("%.0f°", temp)),
		feelsStyle.Render(fmt.Sprintf("%.0f°", feels_like)),
	)

	content := lipgloss.JoinVertical(lipgloss.Center,
		headerRow,
		dataRow,
	)

	fmt.Println(borderStyle.Align(lipgloss.Center).Render(content))
}
