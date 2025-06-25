#!/bin/bash
#
# TIV Demo Script
# Demonstrates the features of Terminal Image Viewer
#

echo "=== TIV Demo - Terminal Image Viewer ==="
echo "Converting images to ASCII art with Unix philosophy"
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
    ./build/tiv -w 40 test.png | head -6
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
    ./build/tiv -w 30 test.png | head -4
    echo
    echo "   Enhanced contrast (1.8):"
    ./build/tiv -c 1.8 -w 30 test.png | head -4
    echo
fi

echo "5. Unix philosophy in action:"
echo "   All features work with pipes and redirection:"
echo "   • cat image.jpg | tiv"
echo "   • tiv image.png > ascii.txt"
echo "   • tiv image.jpg | head -10"
echo "   • find . -name '*.jpg' | head -1 | xargs tiv"
echo

echo "=== TIV Features ==="
echo "✓ 13-character ASCII density range"
echo "✓ Box filter sampling for quality"  
echo "✓ Proper aspect ratio handling"
echo "✓ Contrast adjustment (-c flag)"
echo "✓ Unix philosophy compliant"
echo "✓ Cross-platform (Linux, macOS, Windows)"
echo
echo "Try real images:"
echo "  curl -s 'https://picsum.photos/300' | ./build/tiv -c 1.4 -w 60"
echo "  ./build/tiv -w 80 -c 1.2 your-photo.jpg | less"

# Clean up test file if we created it
if [ -f "test.png" ] && command -v convert &> /dev/null; then
    echo
    echo "Cleaning up test.png..."
    rm -f test.png
fi 