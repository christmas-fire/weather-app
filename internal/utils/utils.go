package utils

import "unicode"

// GetColorForTemperature returns a hex color code based on the temperature value.
// The color codes represent a gradient from cold to hot temperatures.
func GetColorForTemperature(temp float64) string {
	switch {
	case temp <= -30:
		return "#191970"
	case temp <= -15:
		return "#0000CD"
	case temp <= 0:
		return "#248AFD"
	case temp <= 10:
		return "#00FFFF"
	case temp <= 20:
		return "#90EE90"
	case temp <= 29:
		return "#FFA500"
	case temp <= 39:
		return "#FF4500"
	default:
		return "#DC143C" // >= 40
	}
}

// CapitalizeFirst capitalizes the first letter of the given string.
// It correctly handles Unicode characters.
func CapitalizeFirst(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
