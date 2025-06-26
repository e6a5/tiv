package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

// Version is set by build flags
var Version = "development"

func main() {
	var config Config
	var showVersion bool
	var colorMode string
	
	// Define CLI flags
	flag.IntVar(&config.Width, "w", 80, "Output width in characters")
	flag.IntVar(&config.Width, "width", 80, "Output width in characters")
	flag.IntVar(&config.Height, "h", 0, "Output height in characters (auto-calculated if 0)")
	flag.IntVar(&config.Height, "height", 0, "Output height in characters (auto-calculated if 0)")
	flag.BoolVar(&config.Invert, "i", false, "Invert brightness levels")
	flag.BoolVar(&config.Invert, "invert", false, "Invert brightness levels")
	flag.Float64Var(&config.Contrast, "c", 1.0, "Contrast adjustment (0.5-2.0, default 1.0)")
	flag.Float64Var(&config.Contrast, "contrast", 1.0, "Contrast adjustment (0.5-2.0, default 1.0)")
	flag.BoolVar(&config.UseBlocks, "b", false, "Use Unicode block characters for higher resolution")
	flag.BoolVar(&config.UseBlocks, "blocks", false, "Use Unicode block characters for higher resolution")
	flag.BoolVar(&config.Dither, "d", false, "Apply Floyd-Steinberg dithering for smoother gradients")
	flag.BoolVar(&config.Dither, "dither", false, "Apply Floyd-Steinberg dithering for smoother gradients")
	flag.StringVar(&colorMode, "color", "", "Enable color output: '256' for 256-color, '24bit' for truecolor")
	flag.BoolVar(&config.Preview, "p", false, "Show original image inline (instead of ASCII)")
	flag.BoolVar(&config.Preview, "preview", false, "Show original image inline (instead of ASCII)")
	flag.StringVar(&config.PreviewMode, "preview-mode", "auto", "Preview mode: 'auto', 'terminal', or 'system'")
	flag.BoolVar(&config.NoSplit, "no-split", false, "Disable split view (show ASCII only)")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [image_file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Show images side-by-side with ASCII art. Reads from stdin if no file specified.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s image.jpg                           # Split view: image + ASCII\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 40 image.png                     # Split view with custom width\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -no-split image.jpg                 # ASCII only (classic mode)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -c 1.5 image.jpg                    # Split view with contrast\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -b image.jpg                        # Split view with blocks\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -d image.jpg                        # Split view with dithering\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -color 24bit image.jpg              # Split view with color\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -b -d -color 24bit -c 1.3 image.jpg # Ultimate split view\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -p image.jpg                        # Image only (no ASCII)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -p -preview-mode system image.jpg   # External viewer\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  cat image.jpg | %s                     # Pipe mode (ASCII only)\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -no-split image.jpg > ascii.txt     # Save ASCII to file\n", os.Args[0])
	}
	
	flag.Parse()
	
	// Handle version flag
	if showVersion {
		fmt.Printf("TIV (Terminal Image Viewer) %s\n", Version)
		fmt.Println("Convert images to ASCII art - Unix philosophy compliant")
		fmt.Println("https://github.com/e6a5/tiv")
		return
	}
	
	// Parse and validate inputs
	reader, filename, err := parseInputSource()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", friendlyError(err, "parsing input"))
		os.Exit(1)
	}
	
	// Validate configuration
	if err := validateConfig(&config); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	
	// Validate image file if provided
	if filename != "" {
		if err := validateImageFile(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	}
	
	// Validate preview flag usage
	if config.Preview && filename == "" {
		fmt.Fprintf(os.Stderr, "Error: preview flag requires a file argument\n")
		flag.Usage()
		os.Exit(1)
	}
	
	// Parse color mode
	config.Color = parseColorMode(colorMode)
	
	// Determine operation mode and execute
	if filename != "" && config.Preview {
		// Preview mode: show image only
		handlePreviewMode(filename, parsePreviewMode(config.PreviewMode), reader, config)
	} else if filename != "" && !config.NoSplit {
		// Split view mode: image + ASCII (default for files)
		handleSplitViewMode(filename, reader, config)
	} else {
		// ASCII only mode: for stdin or when --no-split is used
		handleASCIIMode(reader, config)
	}
}

// parseInputSource determines input source (file or stdin)
func parseInputSource() (io.Reader, string, error) {
	if flag.NArg() == 0 {
		// Read from stdin - no preview support for stdin
		return os.Stdin, "", nil
	} else if flag.NArg() == 1 {
		// Read from file
		filename := flag.Arg(0)
		file, err := os.Open(filename)
		if err != nil {
			return nil, "", fmt.Errorf("opening file: %w", err)
		}
		return file, filename, nil
	} else {
		return nil, "", fmt.Errorf("too many arguments")
	}
}

// parseColorMode converts string to ColorMode
func parseColorMode(colorMode string) ColorMode {
	switch colorMode {
	case "256":
		return Color256
	case "24bit", "truecolor":
		return Color24bit
	default:
		return ColorNone
	}
}

// handlePreviewMode processes preview-only mode
func handlePreviewMode(filename string, mode PreviewMode, reader io.Reader, config Config) {
	if err := showImagePreview(filename, mode); err != nil {
		// Fallback to ASCII if preview fails
		fmt.Fprintf(os.Stderr, "Image preview not supported, showing ASCII conversion...\n")
		handleASCIIMode(reader, config)
	}
}

// handleSplitViewMode processes split view mode (default for files)
func handleSplitViewMode(filename string, reader io.Reader, config Config) {
	// Adjust config for split view dimensions
	termWidth, _ := getTerminalSize()
	splitConfig := config
	if splitConfig.Width == 80 { // Only adjust default width
		splitConfig.Width = termWidth/2 - 1
	}
	
	// Generate ASCII for right side
	asciiArt, err := generateASCII(reader, splitConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating ASCII: %v\n", err)
		os.Exit(1)
	}
	
	// Show split view
	mode := parsePreviewMode(splitConfig.PreviewMode)
	if err := showSplitView(filename, asciiArt, mode); err != nil {
		fmt.Fprintf(os.Stderr, "Error showing split view: %v\n", err)
		os.Exit(1)
	}
}

// handleASCIIMode processes ASCII-only mode
func handleASCIIMode(reader io.Reader, config Config) {
	if err := processImage(reader, config); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing image: %v\n", err)
		os.Exit(1)
	}
}

// generateASCII reads an image and generates ASCII art string
func generateASCII(reader io.Reader, config Config) (string, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return "", friendlyError(fmt.Errorf("failed to decode image: %w", err), "image decoding")
	}
	
	// Validate image dimensions
	if err := validateImageDimensions(img, "input image"); err != nil {
		return "", err
	}
	
	// Use memory-optimized processing for large images
	processor := NewChunkedProcessor(config)
	result, err := processor.processLargeImageOptimized(img)
	if err != nil {
		return "", friendlyError(err, "image processing")
	}
	
	return result, nil
}

// processImage reads an image and converts it to ASCII
func processImage(reader io.Reader, config Config) error {
	result, err := generateASCII(reader, config)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

func init() {
	// Register image formats
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

 