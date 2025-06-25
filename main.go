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
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [image_file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Convert images to ASCII art. Reads from stdin if no file specified.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s image.jpg\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 40 image.png\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -c 1.5 image.jpg                    # Increase contrast\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -b image.jpg                        # High resolution blocks\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -d image.jpg                        # Smooth dithering\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -color 256 image.jpg                # 256-color output\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -color 24bit image.jpg              # Full truecolor\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -b -d -c 1.3 image.jpg              # Best quality B&W\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -b -d -color 24bit -c 1.3 image.jpg # Ultimate quality\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  cat image.jpg | %s\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s image.jpg > ascii.txt\n", os.Args[0])
	}
	
	flag.Parse()
	
	// Handle version flag
	if showVersion {
		fmt.Printf("TIV (Terminal Image Viewer) %s\n", Version)
		fmt.Println("Convert images to ASCII art - Unix philosophy compliant")
		fmt.Println("https://github.com/e6a5/tiv")
		return
	}
	
	var reader io.Reader
	
	// Determine input source
	if flag.NArg() == 0 {
		// Read from stdin
		reader = os.Stdin
	} else if flag.NArg() == 1 {
		// Read from file
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	} else {
		fmt.Fprintf(os.Stderr, "Error: too many arguments\n")
		flag.Usage()
		os.Exit(1)
	}
	
	// Parse color mode
	switch colorMode {
	case "":
		config.Color = ColorNone
	case "256":
		config.Color = Color256
	case "24bit", "truecolor":
		config.Color = Color24bit
	default:
		fmt.Fprintf(os.Stderr, "Error: invalid color mode '%s'. Use '256' or '24bit'\n", colorMode)
		os.Exit(1)
	}
	
	// Process the image
	if err := processImage(reader, config); err != nil {
		fmt.Fprintf(os.Stderr, "Error processing image: %v\n", err)
		os.Exit(1)
	}
}

// processImage reads an image and converts it to ASCII
func processImage(reader io.Reader, config Config) error {
	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		return fmt.Errorf("failed to decode image: %w", err)
	}
	
	// Convert to ASCII or blocks and output
	var result string
	if config.Dither {
		result = imageToArtWithDithering(img, config)
	} else if config.UseBlocks {
		result = imageToBlocks(img, config)
	} else {
		result = imageToASCII(img, config)
	}
	fmt.Print(result)
	
	return nil
}

func init() {
	// Register image formats
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
}

 