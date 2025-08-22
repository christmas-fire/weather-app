package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/christmas-fire/weather-app/config"
	"github.com/christmas-fire/weather-app/internal/models"
	"github.com/christmas-fire/weather-app/internal/out"
)

type weatherDataMsg models.WeatherData

type errMsg struct {
	err error
}

func (e errMsg) Error() string {
	return e.err.Error()
}

type model struct {
	spinner spinner.Model
	weather *models.WeatherData
	err     error
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#248AFD"))
	return model{
		spinner: s,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		getWeatherData,
		m.spinner.Tick,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m, tea.Quit

	case weatherDataMsg:
		weather := models.WeatherData(msg)
		m.weather = &weather
		return m, tea.Quit

	case errMsg:
		m.err = msg
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m model) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nerror: %v\n\n", m.err)
	}

	if m.weather != nil {
		return ""
	}

	return fmt.Sprintf("\n   %s Loading...\n\n", m.spinner.View())
}

func getWeatherData() tea.Msg {
	if err := config.LoadConfig(); err != nil {
		return errMsg{fmt.Errorf("error init config: %s", err.Error())}
	}

	apiKey := os.Getenv("API_KEY")
	city := os.Getenv("CITY")

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	time.Sleep(1 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		return errMsg{err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return errMsg{fmt.Errorf("API error: %s (code: %d)", string(body), resp.StatusCode)}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errMsg{err}
	}

	var weatherData models.WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return errMsg{err}
	}

	return weatherDataMsg(weatherData)
}

func main() {
	p := tea.NewProgram(initialModel())

	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	if finalModel, ok := m.(model); ok {
		out.PrintWeatherData(*finalModel.weather)
	}
}
