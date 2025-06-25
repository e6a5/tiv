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

// ASCII characters ordered by brightness (darkest to lightest)
// Using more characters for better density representation
const asciiChars = " .':;!>*+%S#@"

// Config holds the CLI options
type Config struct {
	Width    int
	Height   int
	Invert   bool
	Contrast float64
}

func main() {
	var config Config
	var showVersion bool
	
	// Define CLI flags
	flag.IntVar(&config.Width, "w", 80, "Output width in characters")
	flag.IntVar(&config.Width, "width", 80, "Output width in characters")
	flag.IntVar(&config.Height, "h", 0, "Output height in characters (auto-calculated if 0)")
	flag.IntVar(&config.Height, "height", 0, "Output height in characters (auto-calculated if 0)")
	flag.BoolVar(&config.Invert, "i", false, "Invert brightness levels")
	flag.BoolVar(&config.Invert, "invert", false, "Invert brightness levels")
	flag.Float64Var(&config.Contrast, "c", 1.0, "Contrast adjustment (0.5-2.0, default 1.0)")
	flag.Float64Var(&config.Contrast, "contrast", 1.0, "Contrast adjustment (0.5-2.0, default 1.0)")
	flag.BoolVar(&showVersion, "version", false, "Show version information")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [image_file]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Convert images to ASCII art. Reads from stdin if no file specified.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  %s image.jpg\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -w 40 image.png\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s -c 1.5 image.jpg        # Increase contrast\n", os.Args[0])
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
	
	// Convert to ASCII and output
	ascii := imageToASCII(img, config)
	fmt.Print(ascii)
	
	return nil
}

// imageToASCII converts an image to ASCII art
func imageToASCII(img image.Image, config Config) string {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	
	// Calculate output dimensions
	outWidth := config.Width
	outHeight := config.Height
	
	// Auto-calculate height if not specified (maintain aspect ratio)
	if outHeight == 0 {
		// Adjust for character aspect ratio (characters are typically taller than wide)
		aspectRatio := float64(height) / float64(width)
		outHeight = int(float64(outWidth) * aspectRatio * 0.43) // Better character aspect ratio
	}
	
	var result string
	
	for y := 0; y < outHeight; y++ {
		for x := 0; x < outWidth; x++ {
			// Sample multiple pixels for better quality (simple box filter)
			var totalR, totalG, totalB uint64
			samples := 0
			
			// Calculate source region to sample
			startX := float64(x) * float64(width) / float64(outWidth)
			endX := float64(x+1) * float64(width) / float64(outWidth)
			startY := float64(y) * float64(height) / float64(outHeight)
			endY := float64(y+1) * float64(height) / float64(outHeight)
			
			// Sample the region (at least the center pixel)
			minX := int(startX)
			maxX := int(endX)
			minY := int(startY)
			maxY := int(endY)
			
			if minX == maxX {
				maxX = minX + 1
			}
			if minY == maxY {
				maxY = minY + 1
			}
			
			// Clamp to image bounds
			if minX < 0 { minX = 0 }
			if maxX >= width { maxX = width - 1 }
			if minY < 0 { minY = 0 }
			if maxY >= height { maxY = height - 1 }
			
			for sy := minY; sy <= maxY; sy++ {
				for sx := minX; sx <= maxX; sx++ {
					r, g, b, _ := img.At(sx, sy).RGBA()
					totalR += uint64(r)
					totalG += uint64(g)
					totalB += uint64(b)
					samples++
				}
			}
			
			// Average the samples
			avgR := totalR / uint64(samples)
			avgG := totalG / uint64(samples)
			avgB := totalB / uint64(samples)
			
			// Convert to grayscale (luminance formula)
			// Using standard luminance weights: 0.299*R + 0.587*G + 0.114*B
			gray := (299*avgR + 587*avgG + 114*avgB) / 1000
			
			// Normalize to 0-255 range (RGBA values are 0-65535)
			gray = gray >> 8
			
			// Apply contrast adjustment
			adjustedGray := applyContrast(int(gray), config.Contrast)
			
			// Convert to ASCII character
			char := grayToASCII(adjustedGray, config.Invert)
			result += string(char)
		}
		result += "\n"
	}
	
	return result
}

// grayToASCII converts a grayscale value (0-255) to an ASCII character
func grayToASCII(gray int, invert bool) byte {
	if invert {
		gray = 255 - gray
	}
	
	// Map gray value to ASCII character index
	index := gray * (len(asciiChars) - 1) / 255
	if index >= len(asciiChars) {
		index = len(asciiChars) - 1
	}
	
	return asciiChars[index]
}

// applyContrast adjusts the contrast of a grayscale value
func applyContrast(gray int, contrast float64) int {
	if contrast == 1.0 {
		return gray
	}
	
	// Apply contrast: ((gray/255 - 0.5) * contrast + 0.5) * 255
	normalized := float64(gray) / 255.0
	adjusted := (normalized-0.5)*contrast + 0.5
	
	// Clamp to 0-255 range
	if adjusted < 0 {
		adjusted = 0
	} else if adjusted > 1 {
		adjusted = 1
	}
	
	return int(adjusted * 255)
}

func init() {
	// Register image formats
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
} 