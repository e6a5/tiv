package main

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	
	// Extended format support
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

// getSupportedFormats returns a list of supported image formats
func getSupportedFormats() []string {
	return []string{
		"PNG (.png)",
		"JPEG (.jpg, .jpeg)", 
		"GIF (.gif)",
		"WebP (.webp)",
		"TIFF (.tiff, .tif)",
		"BMP (.bmp)",
	}
}

// Note: Helper functions for future use when adding batch processing or format detection features
// These will be used in Phase 2 implementation for enhanced functionality 