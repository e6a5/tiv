#!/bin/bash
#
# TIV Demo Script
# Demonstrates the features of Terminal Image Viewer
#

echo "=== TIV Demo - Terminal Image Viewer ==="
echo "Professional ASCII art with Unicode blocks & Floyd-Steinberg dithering"
echo

# Create a simple test gradient if no image provided
if [ ! -f "test.png" ] && command -v convert &> /dev/null; then
    echo "Creating test gradient image..."
    convert -size 200x100 gradient:black-white test.png
    echo "✓ Created test.png"
    echo
fi

echo "1. Enhanced ASCII character set for smooth gradients:"
echo "   Character set: ' .':;!>*+%S#@' (13 density levels)"
echo

if [ -f "test.png" ]; then
    echo "   Gradient demonstration:"
    ./tiv -w 40 test.png | head -6
    echo
fi

echo "2. Improved sampling algorithm:"
echo "   Before: Single pixel sampling (aliasing issues)"  
echo "   After: Box filter sampling (smoother downscaling)"
echo

echo "3. Better aspect ratio (0.43 vs 0.5):"
echo "   More accurate character proportions"
echo

echo "4. Contrast control for better visibility:"
if [ -f "test.png" ]; then
    echo "   Normal contrast (1.0):"
    ./tiv -w 30 test.png | head -4
    echo
    echo "   Enhanced contrast (1.8):"
    ./tiv -c 1.8 -w 30 test.png | head -4
    echo
fi

echo "5. Unix philosophy in action:"
echo "   All features work with pipes and redirection:"
echo "   • cat image.jpg | tiv"
echo "   • tiv image.png > ascii.txt"
echo "   • tiv image.jpg | head -10"
echo "   • find . -name '*.jpg' | head -1 | xargs tiv"
echo

echo "=== 🚀 TIV Features - Market Leading Quality ==="
echo "✓ ANSI Color Support (256-color & 24-bit truecolor)"
echo "✓ Unicode block characters (2x resolution vs ASCII)"
echo "✓ Floyd-Steinberg dithering (professional gradients)"
echo "✓ 13-character ASCII density range"
echo "✓ Box filter sampling for quality"  
echo "✓ Contrast adjustment (-c flag)"
echo "✓ Unix philosophy compliant (pipeable & composable)"
echo "✓ Cross-platform (Linux, macOS, Windows)"
echo
echo "🎯 Quality modes:"
echo "  Standard:       ./tiv image.jpg"
echo "  High-res:       ./tiv -b image.jpg"
echo "  Smooth:         ./tiv -d image.jpg" 
echo "  Color 256:      ./tiv -color 256 image.jpg"
echo "  Color 24-bit:   ./tiv -color 24bit image.jpg"
echo "  🌟 Ultimate:    ./tiv -b -d -color 24bit -c 1.3 image.jpg"
echo
echo "🌈 Color modes:"
echo "  B&W ASCII:      ./tiv image.jpg"
echo "  256-color:      ./tiv -color 256 image.jpg"
echo "  24-bit color:   ./tiv -color 24bit image.jpg"
echo "  Color + blocks: ./tiv -b -color 24bit image.jpg"
echo "  Full quality:   ./tiv -b -d -color 24bit -c 1.3 image.jpg"
echo
echo "🔗 Unix pipes:"
echo "  curl -s 'https://picsum.photos/300' | ./tiv -b -d -color 24bit -w 60"
echo "  ./tiv -b -d -color 24bit your-photo.jpg | less"

# Clean up test file if we created it
if [ -f "test.png" ] && command -v convert &> /dev/null; then
    echo
    echo "Cleaning up test.png..."
    rm -f test.png
fi 