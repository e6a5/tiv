package main

// ASCII characters ordered by brightness (darkest to lightest)
// Using more characters for better density representation
const asciiChars = " .':;!>*+%S#@"

// Note: blockChars was removed as it was unused.
// The project uses halfBlocks for Unicode block rendering.

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