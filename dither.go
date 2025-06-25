package main

import "image"

// imageToArtWithDithering converts an image to ASCII/blocks with Floyd-Steinberg dithering
func imageToArtWithDithering(img image.Image, config Config) string {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	
	// Calculate output dimensions
	outWidth := config.Width
	outHeight := config.Height
	
	if outHeight == 0 {
		aspectRatio := float64(height) / float64(width)
		outHeight = int(float64(outWidth) * aspectRatio * 0.43)
	}
	
	// Create buffers for dithering
	grayBuffer := make([][]float64, outHeight)
	var colorBuffer [][][3]uint8 // RGB color buffer for color mode
	
	for y := range grayBuffer {
		grayBuffer[y] = make([]float64, outWidth)
	}
	
	if config.Color != ColorNone {
		colorBuffer = make([][][3]uint8, outHeight)
		for y := range colorBuffer {
			colorBuffer[y] = make([][3]uint8, outWidth)
		}
	}
	
	// First pass: fill the buffers with values
	for y := 0; y < outHeight; y++ {
		for x := 0; x < outWidth; x++ {
			// Sample the region for this character
			startX := float64(x) * float64(width) / float64(outWidth)
			endX := float64(x+1) * float64(width) / float64(outWidth)
			startY := float64(y) * float64(height) / float64(outHeight)
			endY := float64(y+1) * float64(height) / float64(outHeight)
			
			minX, maxX, minY, maxY := int(startX), int(endX), int(startY), int(endY)
			
			gray := sampleRegion(img, minX, minY, maxX, maxY)
			adjustedGray := applyContrast(gray, config.Contrast)
			grayBuffer[y][x] = float64(adjustedGray)
			
			// Store color information if needed
			if config.Color != ColorNone {
				r, g, b := sampleRegionColor(img, minX, minY, maxX, maxY)
				colorBuffer[y][x] = [3]uint8{r, g, b}
			}
		}
	}
	
	// Second pass: apply Floyd-Steinberg dithering
	var result string
	for y := 0; y < outHeight; y++ {
		for x := 0; x < outWidth; x++ {
			// Get current pixel value
			oldPixel := grayBuffer[y][x]
			
			// Find closest character and its gray value
			var newPixel float64
			var char string
			
			if config.UseBlocks {
				char, newPixel = findClosestBlock(oldPixel, config.Invert)
			} else {
				char, newPixel = findClosestASCII(oldPixel, config.Invert)
			}
			
			// Apply color if enabled (no dithering on color, just use original)
			if config.Color != ColorNone {
				rgb := colorBuffer[y][x]
				char = colorizeChar(char, rgb[0], rgb[1], rgb[2], config)
			}
			
			result += char
			
			// Calculate quantization error
			error := oldPixel - newPixel
			
			// Distribute error to neighboring pixels (Floyd-Steinberg pattern)
			// X  7/16
			// 3/16 5/16 1/16
			if x+1 < outWidth {
				grayBuffer[y][x+1] += error * 7.0/16.0
			}
			if y+1 < outHeight {
				if x > 0 {
					grayBuffer[y+1][x-1] += error * 3.0/16.0
				}
				grayBuffer[y+1][x] += error * 5.0/16.0
				if x+1 < outWidth {
					grayBuffer[y+1][x+1] += error * 1.0/16.0
				}
			}
		}
		result += "\n"
	}
	
	return result
}

// findClosestASCII finds the closest ASCII character for a grayscale value
func findClosestASCII(gray float64, invert bool) (string, float64) {
	if invert {
		gray = 255.0 - gray
	}
	
	// Clamp gray value
	if gray < 0 { gray = 0 }
	if gray > 255 { gray = 255 }
	
	runes := []rune(asciiChars)
	index := int(gray * float64(len(runes)-1) / 255.0)
	
	if index >= len(runes) {
		index = len(runes) - 1
	}
	if index < 0 {
		index = 0
	}
	
	// Return character and its effective gray value
	effectiveGray := float64(index) * 255.0 / float64(len(runes)-1)
	if invert {
		effectiveGray = 255.0 - effectiveGray
	}
	
	return string(runes[index]), effectiveGray
}

// findClosestBlock finds the closest Unicode block character for a grayscale value
func findClosestBlock(gray float64, invert bool) (string, float64) {
	if invert {
		gray = 255.0 - gray
	}
	
	// Clamp gray value
	if gray < 0 { gray = 0 }
	if gray > 255 { gray = 255 }
	
	runes := []rune(halfBlocks)
	index := int(gray * float64(len(runes)-1) / 255.0)
	
	if index >= len(runes) {
		index = len(runes) - 1
	}
	if index < 0 {
		index = 0
	}
	
	// Return character and its effective gray value
	effectiveGray := float64(index) * 255.0 / float64(len(runes)-1)
	if invert {
		effectiveGray = 255.0 - effectiveGray
	}
	
	return string(runes[index]), effectiveGray
} 