package main

import "image"

// sampleRegion samples a rectangular region and returns average grayscale value
func sampleRegion(img image.Image, minX, minY, maxX, maxY int) int {
	bounds := img.Bounds()
	
	// Clamp to image bounds
	if minX < bounds.Min.X { minX = bounds.Min.X }
	if maxX >= bounds.Max.X { maxX = bounds.Max.X - 1 }
	if minY < bounds.Min.Y { minY = bounds.Min.Y }
	if maxY >= bounds.Max.Y { maxY = bounds.Max.Y - 1 }
	
	if minX >= maxX { maxX = minX + 1 }
	if minY >= maxY { maxY = minY + 1 }
	
	var totalR, totalG, totalB uint64
	samples := 0
	
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			totalR += uint64(r)
			totalG += uint64(g)
			totalB += uint64(b)
			samples++
		}
	}
	
	if samples == 0 {
		return 0
	}
	
	// Average the samples
	avgR := totalR / uint64(samples)
	avgG := totalG / uint64(samples)
	avgB := totalB / uint64(samples)
	
	// Convert to grayscale
	gray := (299*avgR + 587*avgG + 114*avgB) / 1000
	return int(gray >> 8) // Normalize to 0-255
}

// sampleRegionColor samples a region and returns average RGB values
func sampleRegionColor(img image.Image, minX, minY, maxX, maxY int) (uint8, uint8, uint8) {
	var totalR, totalG, totalB uint64
	samples := 0
	
	for sy := minY; sy <= maxY; sy++ {
		for sx := minX; sx <= maxX; sx++ {
			r, g, b, _ := img.At(sx, sy).RGBA()
			totalR += uint64(r)
			totalG += uint64(g)
			totalB += uint64(b)
			samples++
		}
	}
	
	if samples == 0 {
		return 0, 0, 0
	}
	
	// Average and convert from 16-bit to 8-bit
	avgR := uint8((totalR / uint64(samples)) >> 8)
	avgG := uint8((totalG / uint64(samples)) >> 8)
	avgB := uint8((totalB / uint64(samples)) >> 8)
	
	return avgR, avgG, avgB
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