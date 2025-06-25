package main

import "fmt"

// colorizeChar wraps a character with ANSI color codes
func colorizeChar(char string, r, g, b uint8, config Config) string {
	if config.Color == ColorNone {
		return char
	}
	
	var colorCode string
	switch config.Color {
	case Color256:
		colorCode = fmt.Sprintf("\033[38;5;%dm", rgbTo256Color(r, g, b))
	case Color24bit:
		colorCode = fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	}
	
	return colorCode + char + "\033[0m"
}

// rgbTo256Color converts RGB values to the closest 256-color palette index
func rgbTo256Color(r, g, b uint8) int {
	// For colors 16-231: 6x6x6 color cube
	if r == g && g == b {
		// Grayscale colors 232-255
		if r < 8 {
			return 16
		}
		if r > 248 {
			return 231
		}
		return 232 + (int(r)-8)/10
	}
	
	// Map to 6x6x6 color cube (colors 16-231)
	rIndex := int(r) * 5 / 255
	gIndex := int(g) * 5 / 255
	bIndex := int(b) * 5 / 255
	
	return 16 + 36*rIndex + 6*gIndex + bIndex
} 