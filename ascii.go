package main

import "image"

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
			
			// Get color information if needed
			var r, g, b uint8
			if config.Color != ColorNone {
				r, g, b = sampleRegionColor(img, minX, minY, maxX, maxY)
			}
			
			// Sample for grayscale
			gray := sampleRegion(img, minX, minY, maxX, maxY)
			
			// Apply contrast adjustment
			adjustedGray := applyContrast(gray, config.Contrast)
			
			// Convert to ASCII character
			char := grayToASCII(adjustedGray, config.Invert)
			charStr := string(char)
			
			// Apply color if enabled
			if config.Color != ColorNone {
				charStr = colorizeChar(charStr, r, g, b, config)
			}
			
			result += charStr
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