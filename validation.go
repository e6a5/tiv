package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"strings"
)

// ValidationError represents a validation error with context
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error for %s: %s (got: %v)", e.Field, e.Message, e.Value)
}

// ImageError represents image-related errors with helpful messages
type ImageError struct {
	Type     string
	Filename string
	Err      error
}

func (e ImageError) Error() string {
	return fmt.Sprintf("%s: %s (%v)", e.Type, e.Filename, e.Err)
}

// validateConfig validates all configuration parameters
func validateConfig(config *Config) error {
	// Validate width
	if config.Width < 1 {
		return ValidationError{
			Field:   "width",
			Value:   config.Width,
			Message: "must be at least 1",
		}
	}
	if config.Width > 1000 {
		return ValidationError{
			Field:   "width",
			Value:   config.Width,
			Message: "must be 1000 or less to prevent performance issues",
		}
	}

	// Validate height
	if config.Height < 0 {
		return ValidationError{
			Field:   "height",
			Value:   config.Height,
			Message: "must be 0 (auto) or positive",
		}
	}
	if config.Height > 1000 {
		return ValidationError{
			Field:   "height",
			Value:   config.Height,
			Message: "must be 1000 or less to prevent performance issues",
		}
	}

	// Validate contrast
	if config.Contrast < 0.1 {
		return ValidationError{
			Field:   "contrast",
			Value:   config.Contrast,
			Message: "must be at least 0.1",
		}
	}
	if config.Contrast > 5.0 {
		return ValidationError{
			Field:   "contrast",
			Value:   config.Contrast,
			Message: "must be 5.0 or less",
		}
	}

	// Validate preview mode
	validPreviewModes := []string{"auto", "terminal", "system"}
	isValidMode := false
	for _, mode := range validPreviewModes {
		if config.PreviewMode == mode {
			isValidMode = true
			break
		}
	}
	if !isValidMode {
		return ValidationError{
			Field:   "preview-mode",
			Value:   config.PreviewMode,
			Message: fmt.Sprintf("must be one of: %s", strings.Join(validPreviewModes, ", ")),
		}
	}

	return nil
}

// validateImageFile validates image file existence and format
func validateImageFile(filename string) error {
	if filename == "" {
		return ValidationError{
			Field:   "filename",
			Value:   filename,
			Message: "filename cannot be empty",
		}
	}

	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return ImageError{
			Type:     "File not found",
			Filename: filename,
			Err:      fmt.Errorf("check the path and try again"),
		}
	}

	// Check file extension (basic validation)
	ext := strings.ToLower(filepath.Ext(filename))
	supportedExts := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".tiff", ".tif", ".bmp"}
	isSupported := false
	for _, supportedExt := range supportedExts {
		if ext == supportedExt {
			isSupported = true
			break
		}
	}

	if !isSupported {
		return ImageError{
			Type:     "Unsupported image format",
			Filename: filename,
			Err:      fmt.Errorf("supported formats: %s", strings.Join(getSupportedFormats(), ", ")),
		}
	}

	return nil
}

// validateImageDimensions validates image dimensions for processing
func validateImageDimensions(img image.Image, filename string) error {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	if width < 1 || height < 1 {
		return ImageError{
			Type:     "Invalid image dimensions",
			Filename: filename,
			Err:      fmt.Errorf("image has dimensions %dx%d", width, height),
		}
	}

	// Warn about very large images that might cause memory issues
	const maxDimension = 8000
	const maxPixels = 20_000_000 // 20 megapixels

	if width > maxDimension || height > maxDimension {
		return ImageError{
			Type:     "Image too large",
			Filename: filename,
			Err:      fmt.Errorf("dimensions %dx%d exceed maximum %d on any side", width, height, maxDimension),
		}
	}

	if width*height > maxPixels {
		return ImageError{
			Type:     "Image too large",
			Filename: filename,
			Err:      fmt.Errorf("image has %d pixels, maximum supported is %d", width*height, maxPixels),
		}
	}

	return nil
}

// friendlyError converts technical errors to user-friendly messages
func friendlyError(err error, context string) error {
	if err == nil {
		return nil
	}

	// Check if it's already a friendly error type
	if _, ok := err.(ValidationError); ok {
		return err
	}
	if _, ok := err.(ImageError); ok {
		return err
	}

	errStr := err.Error()
	switch {
	case strings.Contains(errStr, "unknown format"):
		return fmt.Errorf("unsupported image format - try PNG, JPEG, GIF, WebP, TIFF, or BMP")
	case strings.Contains(errStr, "permission denied"):
		return fmt.Errorf("permission denied - check file permissions and try again")
	case strings.Contains(errStr, "no such file"):
		return fmt.Errorf("file not found - check the path and try again")
	case strings.Contains(errStr, "connection refused"):
		return fmt.Errorf("cannot connect to display - check your terminal settings")
	case strings.Contains(errStr, "decode"):
		return fmt.Errorf("corrupted or invalid image file - try a different image")
	case strings.Contains(errStr, "out of memory"):
		return fmt.Errorf("image too large for available memory - try a smaller image or reduce output dimensions")
	default:
		if context != "" {
			return fmt.Errorf("%s: %v", context, err)
		}
		return err
	}
}

// Note: validateTerminalSize function removed as it's not currently used
// Terminal size validation could be added in future if needed for UX improvements 