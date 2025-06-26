package main

import (
	"fmt"
	"image"
	"image/color"
	"runtime"
	"sync"
)

// MemoryLimits defines memory usage constraints
type MemoryLimits struct {
	MaxPixels      int // Maximum pixels to process at once
	ChunkSize      int // Size of processing chunks
	MaxGoroutines  int // Maximum concurrent goroutines
	BufferSizeKB   int // Buffer size in KB
}

// getMemoryLimits returns appropriate memory limits based on system resources
func getMemoryLimits() MemoryLimits {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	
	// Conservative limits based on available memory
	return MemoryLimits{
		MaxPixels:     10_000_000, // 10 megapixels max per chunk
		ChunkSize:     1000,       // 1000x1000 pixel chunks
		MaxGoroutines: runtime.NumCPU(),
		BufferSizeKB:  1024,       // 1MB buffer
	}
}

// ChunkedProcessor handles large image processing in chunks
type ChunkedProcessor struct {
	limits MemoryLimits
	config Config
}

// NewChunkedProcessor creates a new chunked processor
func NewChunkedProcessor(config Config) *ChunkedProcessor {
	return &ChunkedProcessor{
		limits: getMemoryLimits(),
		config: config,
	}
}

// processLargeImageOptimized processes large images with memory optimization
func (cp *ChunkedProcessor) processLargeImageOptimized(img image.Image) (string, error) {
	bounds := img.Bounds()
	imgWidth, imgHeight := bounds.Dx(), bounds.Dy()
	
	// Check if image needs chunked processing
	totalPixels := imgWidth * imgHeight
	if totalPixels <= cp.limits.MaxPixels {
		// Small enough to process normally
		return cp.processRegularImage(img)
	}
	
	// Large image - use chunked processing
	return cp.processImageInChunks(img)
}

// processRegularImage processes smaller images normally
func (cp *ChunkedProcessor) processRegularImage(img image.Image) (string, error) {
	if cp.config.Dither {
		return imageToArtWithDithering(img, cp.config), nil
	} else if cp.config.UseBlocks {
		return imageToBlocks(img, cp.config), nil
	}
	return imageToASCII(img, cp.config), nil
}

// processImageInChunks processes large images in smaller chunks
func (cp *ChunkedProcessor) processImageInChunks(img image.Image) (string, error) {
	bounds := img.Bounds()
	imgWidth, imgHeight := bounds.Dx(), bounds.Dy()
	
	// Calculate chunk dimensions
	chunkWidth := cp.limits.ChunkSize
	chunkHeight := cp.limits.ChunkSize
	
	// Calculate output dimensions
	outWidth := cp.config.Width
	outHeight := cp.config.Height
	
	if outHeight == 0 {
		aspectRatio := float64(imgHeight) / float64(imgWidth)
		outHeight = int(float64(outWidth) * aspectRatio * 0.43)
	}
	
	// Calculate how many chunks we need
	chunksX := (outWidth + chunkWidth - 1) / chunkWidth
	chunksY := (outHeight + chunkHeight - 1) / chunkHeight
	
	// Process chunks concurrently
	type chunkResult struct {
		x, y int
		data string
		err  error
	}
	
	results := make([][]string, chunksY)
	for i := range results {
		results[i] = make([]string, chunksX)
	}
	
	// Use a semaphore to limit concurrent processing
	semaphore := make(chan struct{}, cp.limits.MaxGoroutines)
	resultChan := make(chan chunkResult, chunksX*chunksY)
	var wg sync.WaitGroup
	
	// Process each chunk
	for chunkY := 0; chunkY < chunksY; chunkY++ {
		for chunkX := 0; chunkX < chunksX; chunkX++ {
			wg.Add(1)
			go func(cx, cy int) {
				defer wg.Done()
				semaphore <- struct{}{} // Acquire semaphore
				defer func() { <-semaphore }() // Release semaphore
				
				result := cp.processChunk(img, cx, cy, chunkWidth, chunkHeight, outWidth, outHeight)
				resultChan <- chunkResult{x: cx, y: cy, data: result.data, err: result.err}
			}(chunkX, chunkY)
		}
	}
	
	// Wait for all chunks to complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	
	// Collect results
	for result := range resultChan {
		if result.err != nil {
			return "", fmt.Errorf("chunk processing error at (%d,%d): %w", result.x, result.y, result.err)
		}
		results[result.y][result.x] = result.data
	}
	
	// Combine chunks into final result
	return cp.combineChunks(results), nil
}

// chunkProcessResult represents the result of processing a single chunk
type chunkProcessResult struct {
	data string
	err  error
}

// processChunk processes a single chunk of the image
func (cp *ChunkedProcessor) processChunk(img image.Image, chunkX, chunkY, chunkWidth, chunkHeight, totalOutWidth, totalOutHeight int) chunkProcessResult {
	// Calculate the region of output this chunk represents
	startOutX := chunkX * chunkWidth
	endOutX := startOutX + chunkWidth
	if endOutX > totalOutWidth {
		endOutX = totalOutWidth
	}
	
	startOutY := chunkY * chunkHeight
	endOutY := startOutY + chunkHeight
	if endOutY > totalOutHeight {
		endOutY = totalOutHeight
	}
	
	actualChunkWidth := endOutX - startOutX
	actualChunkHeight := endOutY - startOutY
	
	if actualChunkWidth <= 0 || actualChunkHeight <= 0 {
		return chunkProcessResult{data: "", err: nil}
	}
	
	// Create a config for this chunk
	chunkConfig := cp.config
	chunkConfig.Width = actualChunkWidth
	chunkConfig.Height = actualChunkHeight
	
	// Create a sub-image for this region
	bounds := img.Bounds()
	imgWidth, imgHeight := bounds.Dx(), bounds.Dy()
	
	// Map output coordinates to image coordinates
	startImgX := int(float64(startOutX) * float64(imgWidth) / float64(totalOutWidth))
	endImgX := int(float64(endOutX) * float64(imgWidth) / float64(totalOutWidth))
	startImgY := int(float64(startOutY) * float64(imgHeight) / float64(totalOutHeight))
	endImgY := int(float64(endOutY) * float64(imgHeight) / float64(totalOutHeight))
	
	// Create cropped image
	cropRect := image.Rect(startImgX, startImgY, endImgX, endImgY)
	croppedImg := &croppedImage{
		img:  img,
		rect: cropRect,
	}
	
	// Process the chunk
	var result string
	if chunkConfig.Dither {
		result = imageToArtWithDithering(croppedImg, chunkConfig)
	} else if chunkConfig.UseBlocks {
		result = imageToBlocks(croppedImg, chunkConfig)
	} else {
		result = imageToASCII(croppedImg, chunkConfig)
	}
	
	return chunkProcessResult{data: result, err: nil}
}

// combineChunks combines processed chunks into the final result
func (cp *ChunkedProcessor) combineChunks(chunks [][]string) string {
	var result string
	
	for _, row := range chunks {
		// For each row of chunks, we need to combine horizontally
		if len(row) == 0 {
			continue
		}
		
		// Split each chunk into lines
		chunkLines := make([][]string, len(row))
		maxLines := 0
		
		for x, chunk := range row {
			if chunk == "" {
				chunkLines[x] = []string{}
				continue
			}
			chunkLines[x] = splitLines(chunk)
			if len(chunkLines[x]) > maxLines {
				maxLines = len(chunkLines[x])
			}
		}
		
		// Combine lines horizontally
		for lineIdx := 0; lineIdx < maxLines; lineIdx++ {
			var line string
			for x := 0; x < len(row); x++ {
				if lineIdx < len(chunkLines[x]) {
					line += chunkLines[x][lineIdx]
				}
			}
			if line != "" {
				result += line + "\n"
			}
		}
	}
	
	return result
}

// splitLines splits a string into lines, handling different line endings
func splitLines(s string) []string {
	if s == "" {
		return []string{}
	}
	
	lines := []string{}
	current := ""
	
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, current)
			current = ""
		} else if r != '\r' { // Skip \r characters
			current += string(r)
		}
	}
	
	// Add final line if it doesn't end with newline
	if current != "" {
		lines = append(lines, current)
	}
	
	return lines
}

// croppedImage represents a cropped view of an image
type croppedImage struct {
	img  image.Image
	rect image.Rectangle
}

// ColorModel implements image.Image interface
func (c *croppedImage) ColorModel() color.Model {
	return c.img.ColorModel()
}

// Bounds implements image.Image interface
func (c *croppedImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, c.rect.Dx(), c.rect.Dy())
}

// At implements image.Image interface
func (c *croppedImage) At(x, y int) color.Color {
	// Map coordinates to original image
	origX := c.rect.Min.X + x
	origY := c.rect.Min.Y + y
	
	// Bounds check
	if !image.Pt(origX, origY).In(c.rect) {
		return c.img.ColorModel().Convert(color.Transparent)
	}
	
	return c.img.At(origX, origY)
}

// Note: forceGC function removed as it was unused in current implementation
// Memory management is handled automatically by the chunked processor 