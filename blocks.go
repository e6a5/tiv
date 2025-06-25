package main

import "image"

// imageToBlocks converts an image to Unicode block characters for higher resolution
func imageToBlocks(img image.Image, config Config) string {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	
	// Calculate output dimensions
	outWidth := config.Width
	outHeight := config.Height
	
	// Auto-calculate height if not specified
	// For block mode, we get better vertical resolution, so adjust ratio
	if outHeight == 0 {
		aspectRatio := float64(height) / float64(width)
		outHeight = int(float64(outWidth) * aspectRatio * 0.5) // Same as ASCII for consistency
	}
	
	var result string
	
	for y := 0; y < outHeight; y++ {
		for x := 0; x < outWidth; x++ {
			// Sample the region for this character
			startX := float64(x) * float64(width) / float64(outWidth)
			endX := float64(x+1) * float64(width) / float64(outWidth)
			startY := float64(y) * float64(height) / float64(outHeight)
			endY := float64(y+1) * float64(height) / float64(outHeight)
			
			minX, maxX, minY, maxY := int(startX), int(endX), int(startY), int(endY)
			
			// Get color information if needed
			var r, g, b uint8
			if config.Color != ColorNone {
				r, g, b = sampleRegionColor(img, minX, minY, maxX, maxY)
			}
			
			// Get average brightness for this cell
			gray := sampleRegion(img, minX, minY, maxX, maxY)
			
			// Apply contrast adjustment
			adjustedGray := applyContrast(gray, config.Contrast)
			
			// Convert to block character
			char := grayToBlock(adjustedGray, config.Invert)
			
			// Apply color if enabled
			if config.Color != ColorNone {
				char = colorizeChar(char, r, g, b, config)
			}
			
			result += char
		}
		result += "\n"
	}
	
	return result
}

// grayToBlock converts a grayscale value to a Unicode block character
func grayToBlock(gray int, invert bool) string {
	if invert {
		gray = 255 - gray
	}
	
	// Convert halfBlocks string to runes for proper Unicode handling
	runes := []rune(halfBlocks)
	
	// Map gray value to block character index
	index := gray * (len(runes) - 1) / 255
	if index >= len(runes) {
		index = len(runes) - 1
	}
	if index < 0 {
		index = 0
	}
	
	return string(runes[index])
} 