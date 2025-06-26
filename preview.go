package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// PreviewMode represents different preview modes
type PreviewMode int

const (
	PreviewAuto PreviewMode = iota
	PreviewTerminal
	PreviewSystem
)

// showTerminalPreview displays an image directly in the terminal using various protocols
func showTerminalPreview(filename string, maxWidth, maxHeight int) error {
	// Try protocols in order of preference
	protocols := []func(string, int, int) error{
		tryKittyProtocol,
		tryITermProtocol, 
		trySixelProtocol,
	}
	
	for _, protocol := range protocols {
		if err := protocol(filename, maxWidth, maxHeight); err == nil {
			return nil
		}
	}
	
	return fmt.Errorf("terminal does not support inline image display")
}

// tryKittyProtocol attempts to display image using Kitty terminal protocol
func tryKittyProtocol(filename string, maxWidth, maxHeight int) error {
	// Simple terminal detection
	if !isKittyCompatible() {
		return fmt.Errorf("not a Kitty-compatible terminal")
	}
	
	encoded, err := encodeImageFile(filename)
	if err != nil {
		return err
	}
	
	// Send Kitty graphics protocol
	fmt.Printf("\033_Ga=T,f=100")
	if maxWidth > 0 {
		fmt.Printf(",c=%d", maxWidth)
	}
	if maxHeight > 0 {
		fmt.Printf(",r=%d", maxHeight)
	}
	fmt.Printf(";%s\033\\\n", encoded)
	
	return nil
}

// tryITermProtocol attempts to display image using iTerm2 inline protocol
func tryITermProtocol(filename string, maxWidth, maxHeight int) error {
	if !isITermCompatible() {
		return fmt.Errorf("not an iTerm2-compatible terminal")
	}
	
	encoded, err := encodeImageFile(filename)
	if err != nil {
		return err
	}
	
	// Send iTerm2 inline image protocol
	fmt.Printf("\033]1337;File=inline=1")
	if maxWidth > 0 {
		fmt.Printf(";width=%dpx", maxWidth*8)
	}
	if maxHeight > 0 {
		fmt.Printf(";height=%dpx", maxHeight*16)
	}
	fmt.Printf(":%s\007\n", encoded)
	
	return nil
}

// trySixelProtocol attempts to display image using Sixel protocol
func trySixelProtocol(filename string, maxWidth, maxHeight int) error {
	if !isSixelCompatible() {
		return fmt.Errorf("terminal does not support sixel")
	}
	
	// Try img2sixel command
	if _, err := exec.LookPath("img2sixel"); err == nil {
		cmd := exec.Command("img2sixel", "-w", fmt.Sprintf("%d", maxWidth*8), filename)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	
	return fmt.Errorf("img2sixel not available")
}

// Terminal compatibility checks - simplified
func isKittyCompatible() bool {
	term := os.Getenv("TERM")
	termProgram := os.Getenv("TERM_PROGRAM")
	
	return term == "xterm-kitty" || 
		   os.Getenv("KITTY_WINDOW_ID") != "" ||
		   strings.Contains(termProgram, "ghostty") ||
		   strings.Contains(termProgram, "konsole")
}

func isITermCompatible() bool {
	termProgram := os.Getenv("TERM_PROGRAM")
	
	compatibleTerms := []string{"iTerm", "WezTerm", "warp", "vscode", "Tabby", "Hyper", "Bobcat"}
	for _, term := range compatibleTerms {
		if strings.Contains(termProgram, term) {
			return true
		}
	}
	return false
}

func isSixelCompatible() bool {
	term := os.Getenv("TERM")
	termProgram := os.Getenv("TERM_PROGRAM")
	
	compatibleTerms := []string{"foot", "WindowsTerminal", "blackbox", "xterm"}
	for _, termType := range compatibleTerms {
		if strings.Contains(term, termType) || strings.Contains(termProgram, termType) {
			return true
		}
	}
	return strings.Contains(term, "sixel")
}

// encodeImageFile reads and base64 encodes an image file
func encodeImageFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()
	
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	
	return base64.StdEncoding.EncodeToString(content), nil
}

// parsePreviewMode converts string to PreviewMode
func parsePreviewMode(mode string) PreviewMode {
	switch mode {
	case "terminal":
		return PreviewTerminal
	case "system":
		return PreviewSystem
	default:
		return PreviewAuto
	}
}

// showImagePreview shows an image preview using the best available method
func showImagePreview(filename string, mode PreviewMode) error {
	width, height := getTerminalSize()
	maxWidth := min(width-10, 80)
	maxHeight := min(height-5, 24)
	
	switch mode {
	case PreviewTerminal:
		return showTerminalPreview(filename, maxWidth, maxHeight)
	case PreviewSystem:
		return openSystemViewer(filename)
	case PreviewAuto:
		if err := showTerminalPreview(filename, maxWidth, maxHeight); err != nil {
			return openSystemViewer(filename)
		}
		return nil
	}
	
	return fmt.Errorf("unknown preview mode")
}

// openSystemViewer opens an image file with the system's default viewer
func openSystemViewer(filename string) error {
	var cmd *exec.Cmd
	
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", filename)
	case "linux":
		// Try common viewers
		viewers := []string{"xdg-open", "eog", "feh", "display"}
		for _, viewer := range viewers {
			if _, err := exec.LookPath(viewer); err == nil {
				cmd = exec.Command(viewer, filename)
				break
			}
		}
		if cmd == nil {
			return fmt.Errorf("no suitable image viewer found")
		}
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", filename)
	default:
		return fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	
	return cmd.Start()
}

// getTerminalSize returns terminal dimensions with fallback
func getTerminalSize() (int, int) {
	width, height := 80, 24
	
	if runtime.GOOS != "windows" {
		cmd := exec.Command("stty", "size")
		cmd.Stdin = os.Stdin
		if output, err := cmd.Output(); err == nil {
			// Ignore error from Sscanf as we have fallback values
			_, _ = fmt.Sscanf(string(output), "%d %d", &height, &width)
		}
	}
	
	return width, height
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// showSplitView displays the original image and ASCII side by side
func showSplitView(filename string, asciiArt string, previewMode PreviewMode) error {
	termWidth, termHeight := getTerminalSize()
	
	// Calculate equal dimensions for both sides
	sideWidth := termWidth/2 - 1
	sideHeight := termHeight - 1
	
	// Clear screen and position cursor
	fmt.Print("\033[2J\033[H")
	
	// Try to show image on left side
	if err := showTerminalPreview(filename, sideWidth, sideHeight); err != nil {
		// Show placeholder box if image preview fails
		showPlaceholder(filename, sideWidth, sideHeight, asciiArt)
	}
	
	// Show ASCII on right side
	showASCIIOnRight(asciiArt, termWidth/2, sideWidth, sideHeight)
	
	// Position cursor at bottom
	fmt.Printf("\033[%d;1H", termHeight)
	
	return nil
}

// showPlaceholder displays a placeholder box when image preview fails
func showPlaceholder(filename string, width, height int, asciiArt string) {
	asciiLines := strings.Split(strings.TrimSpace(asciiArt), "\n")
	boxHeight := min(height, len(asciiLines))
	
	// Draw box
	fmt.Printf("┌%s┐\n", strings.Repeat("─", width-2))
	
	for i := 0; i < boxHeight-2; i++ {
		content := ""
		if i == boxHeight/2-1 {
			content = "Original Image"
		} else if i == boxHeight/2 {
			content = fmt.Sprintf("(%s)", filename)
			if len(content) > width-4 {
				content = content[:width-4] + "..."
			}
		}
		
		padding := (width - 2 - len(content)) / 2
		fmt.Printf("│%s%s%s│\n",
			strings.Repeat(" ", padding),
			content,
			strings.Repeat(" ", width-2-padding-len(content)))
	}
	
	fmt.Printf("└%s┘", strings.Repeat("─", width-2))
}

// showASCIIOnRight displays ASCII art on the right side of split view
func showASCIIOnRight(asciiArt string, startCol, width, height int) {
	asciiLines := strings.Split(strings.TrimSpace(asciiArt), "\n")
	
	for i, line := range asciiLines {
		if i >= height {
			break
		}
		
		// Truncate line if too long
		if len(line) > width {
			line = line[:width]
		}
		
		// Position and print line
		fmt.Printf("\033[%d;%dH%s", i+1, startCol, line)
	}
} 