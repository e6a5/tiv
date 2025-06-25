package main

// ASCII characters ordered by brightness (darkest to lightest)
// Using more characters for better density representation
const asciiChars = " .':;!>*+%S#@"

// Unicode block characters for high resolution mode
// These allow for much higher resolution by using top/bottom halves
const blockChars = " ░▒▓█"

// Unicode half-block characters for even higher resolution
const halfBlocks = " ▁▂▃▄▅▆▇█"

// ColorMode represents different color output modes
type ColorMode int

const (
	ColorNone ColorMode = iota
	Color256              // 256-color mode
	Color24bit            // 24-bit truecolor mode
)

// Config holds the CLI options
type Config struct {
	Width       int
	Height      int
	Invert      bool
	Contrast    float64
	UseBlocks   bool
	Dither      bool
	Color       ColorMode
	Preview     bool
	PreviewMode string
	NoSplit     bool
} 